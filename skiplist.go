// Copyright (c) 2014 Dataence, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package skiplist

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sync"

	"github.com/dataence/compare"
)

var (
	DefaultMaxLevel    int     = 12
	DefaultProbability float32 = 0.25
)

type Skiplist struct {
	// Determining MaxLevel
	// Reference: http://drum.lib.umd.edu/bitstream/1903/544/2/CS-TR-2286.1.pdf - section 2
	//
	// > To get away from magic constants, we say that a fraction p of the nodes with level i pointers
	// > also have level i+1 pointers.
	//
	// ip = inverse of p or 1/p or int(math.Ceil(1/p))
	ip int

	// > Since we can safely cap levels at L(n), we should choose MaxLevel = L(N) (where N is an upper
	// > bound on the number of elements in a skip list). If p = 1/2, using MaxLevel = 32 is appropriate
	// > for data structures containing up to 2^32 elements.
	//
	// Magic formula is L = log base 1/p of N or (1/p)^L = N
	//
	// Given p = 1/4 and L = 12, then (1/(1/4))^12 = 4^12 = 2^24 = 16777216 elements in the skiplist
	maxLevel int

	// The number of levels this list has currently. Likely increase until MaxLevel.
	// level is 0-based, so the bottom level is 0, max level is maxLevel-1
	level int

	// headNode is the first node in the skiplist. The next pointers in headNode always points forward
	// to the next node at the appropriate height. Initially all the next pointers will point to tailNode.
	// All of the prev pointers will remain nil.
	headNode *node

	// Using Search Fingers
	// Reference: http://drum.lib.umd.edu/bitstream/1903/544/2/CS-TR-2286.1.pdf - section 3.1
	// We keep two sets of fingers as search and insert localities are likely different, especially if
	// the insert keys are close to each other
	insertFingers []*node

	// fingers for selecting nodes
	selectFingers []*node

	// Total number of nodes inserted
	count int

	// Comparison function for the node keys.
	// For ascending order - if k1 < k2 return true; else return false
	// For descending order - if k1 > k2 return true; else return false
	compare compare.Comparator

	mutex sync.RWMutex
}

func New(compare compare.Comparator) *Skiplist {
	l := DefaultMaxLevel
	ip := int(math.Ceil(1 / float64(DefaultProbability)))

	return &Skiplist{
		ip:            ip,
		maxLevel:      l,
		insertFingers: make([]*node, l),
		selectFingers: make([]*node, l),
		level:         1,
		count:         0,
		compare:       compare,
		headNode:      newNode(l),
	}
}

func (this *Skiplist) SetCompare(compare compare.Comparator) (err error) {
	if compare == nil {
		return errors.New("skiplist/SetCompare: trying to set comparator to nil")
	}
	this.compare = compare
	return nil
}

func (this *Skiplist) SetMaxLevel(l int) (err error) {
	if l < 1 {
		return errors.New("skiplist/SetCompare: max level must be greater than zero (0)")
	}
	this.maxLevel = l
	return nil
}

func (this *Skiplist) SetProbability(p float32) (err error) {
	if p > 1 {
		p = 1
	}
	this.ip = int(math.Ceil(1 / float64(p)))
	return nil
}

func (this *Skiplist) Close() (err error) {
	return nil
}

func (this *Skiplist) Count() int {
	return this.count
}

func (this *Skiplist) Level() int {
	return this.level
}

// Choose the new node's level, branching with p (1/ip) probability, with no regards to N (size of list)
func (this *Skiplist) newNodeLevel() int {
	h := 1

	for h < this.maxLevel && rand.Intn(this.ip) == 0 {
		h++
	}

	return h
}

func (this *Skiplist) updateSearchFingers(key interface{}, fingers []*node) (err error) {
	startLevel := this.level - 1
	startNode := this.headNode

	if fingers[0] != nil && fingers[0].key != nil {
		if less, err := this.compare(fingers[0].key, key); err != nil {
			return err
		} else if less {
			// Move forward, find the highest level s.t. the next node's key < key
			for l := 1; l < this.level; l++ {
				if fingers[l].next[l] != nil && fingers[l].key != nil {
					// If the next node is not nil and fingers[l].key >= key
					if less, err := this.compare(fingers[l].key, key); err != nil {
						return err
					} else if less == false {
						startLevel = l - 1
						startNode = fingers[l]
						break
					}
				}
			}
		} else {
			//log.Println("inside if else, this.level =", this.level-1)
			// Move backward, find the lowest level s.t. the node's timestamp < t
			for l := 1; l < this.level; l++ {
				//log.Println("inside for loop, level =", l)
				// fingers[l].key < key
				if fingers[l].key != nil {
					if less, err := this.compare(fingers[l].key, key); err != nil {
						return err
					} else if less {
						startLevel = l
						startNode = fingers[l]
						break
					}
				}
			}
		}
	}

	// For each of the skiplist levels, going from the current height to 1, walk the list until
	// we find a node that has a timestamp that's >= the timestamp t, or the end of the list
	// l = level, p = ptr to node during traversal
	for l, p := startLevel, startNode; l >= 0; l-- {
		n := p.next[l]

		for {
			if n == nil {
				// last node on the list
				// go to the next level down, and continue traversing
				//log.Println("n == nil")
				break
			}

			//log.Println("n != nil")
			// If n.key >= key
			if less, err := this.compare(n.key, key); err != nil {
				return err
			} else if less == false {
				// Found the first record that either has the same timestamp or greater at this level
				// go to the next level down, and continue traversing
				//log.Println("nt >= t, nt = ", nt.(int64))
				break
			}
			//log.Println("after compare")

			// Move the pointers forward, p = n, n = n.next
			p, n = n, n.next[l]
		}

		fingers[l] = p
	}

	return nil
}

func (this *Skiplist) Insert(key, value interface{}) (*node, error) {
	if key == nil {
		return nil, errors.New("skiplist/Insert: key is nil")
	}

	if this.compare == nil {
		return nil, errors.New("skiplist/Insert: comparator is not set (== nil)")
	}

	// Create new node
	l := this.newNodeLevel()
	n := newNode(l)
	n.SetKey(key)
	n.SetValue(value)

	this.mutex.Lock()
	defer this.mutex.Unlock()

	//log.Println("this.finger[0] =", this.insertFingers[0])
	// Find the position where we should insert the node by updating the search insertFingers using the key
	// Search insertFingers will be updated with the rightmost element of each level that is left of the element
	// that's greater than or equal to key.
	// In other words, we are inserting the new node to the right of the search insertFingers.
	if err := this.updateSearchFingers(key, this.insertFingers); err != nil {
		return nil, errors.New("skiplist/insert: cannot find insert position, " + err.Error())
	}

	//log.Println("search insertFingers =", this.insertFingers)
	// Raise the level of the skiplist if the new level is higher than the existing list level
	// So for levels higher than the current list level, the previous node is headNode for that level
	if this.level < l {
		for i := this.level; i < l; i++ {
			//log.Println("before ---- ", this.insertFingers)
			this.insertFingers[i] = this.headNode
			//log.Println("after  ---- ", this.insertFingers)
		}
		this.level = l
		//log.Println("new this.level =", l)
		//log.Println("this.insertFingers =", this.insertFingers)
	}

	// Finally insert the node into the skiplist
	for i := 0; i < l; i++ {
		// new node points forward to the previous node's next node
		// previous node's next node points to the new node
		n.next[i], this.insertFingers[i].next[i] = this.insertFingers[i].next[i], n
	}

	// Adding to the count
	this.count++

	return n, nil
}

// Select a list of nodes that match the key. The results are stored in the array pointed to by results
func (this *Skiplist) Select(key interface{}) (iter *Iterator, err error) {
	return this.SelectRange(key, key)
}

func (this *Skiplist) SelectRange(key1, key2 interface{}) (iter *Iterator, err error) {
	if key1 == nil || key2 == nil {
		return nil, errors.New("skiplist/SelectRange: key1 or key2 is nil")
	}

	if reflect.TypeOf(key1) != reflect.TypeOf(key2) {
		return nil, fmt.Errorf("skiplist/SelectRange: k1.(%s) and k2.(%s) have different types",
			reflect.TypeOf(key1).Name(), reflect.TypeOf(key2).Name())
	}

	if this.compare == nil {
		return nil, errors.New("skiplist/SelectRange: comparator is not set (== nil)")
	}

	this.mutex.RLock()
	defer this.mutex.RUnlock()

	// Walk the levels and nodes until we find the node at the lowest level (0) that the comparator returns false
	// E.g., if comparator is BuiltinLessThan, then we find the node at the lowest level s.t. node.key < key
	// Then we walk from there to find all the nodes that have node.key == key
	// We keep track of the last touched nodes at each level as selectFingers, and then we re-use the selectFingers
	// so that we can get O(log k) where k is the distance between last searched key and current search key
	// -- ok, so all this is done by updateSearchFingers

	if err = this.updateSearchFingers(key1, this.selectFingers); err != nil {
		return nil, errors.New("skiplist/SelectRange: error selecting nodes, " + err.Error())
	}

	iter = newIterator()
	var res bool
	for p := this.selectFingers[0].next[0]; p != nil; p = p.next[0] {
		pk := p.GetKey()
		if res, err = this.compare(pk, key2); err != nil {
			// If there's error in comparing the keys, then return err
			return nil, errors.New("skiplist/SelectRange: error comparing keys; " + err.Error())
		} else if res || reflect.DeepEqual(pk, key2) {
			iter.buf = append(iter.buf, p)
			iter.count++
		} else {
			// Otherwise if the p.key is "after" key, after could mean greater or less, depending
			// on the comparator, then we know we are done
			break
		}
	}

	return iter, nil
}

func (this *Skiplist) Delete(key interface{}) (iter *Iterator, err error) {
	return this.DeleteRange(key, key)
}

func (this *Skiplist) DeleteRange(key1, key2 interface{}) (iter *Iterator, err error) {
	if key1 == nil || key2 == nil {
		return nil, errors.New("skiplist/DeleteRange: key1 or key2 is nil")
	}

	if reflect.TypeOf(key1) != reflect.TypeOf(key2) {
		return nil, fmt.Errorf("skiplist/DeleteRange: k1.(%s) and k2.(%s) have different types",
			reflect.TypeOf(key1).Name(), reflect.TypeOf(key2).Name())
	}

	if this.compare == nil {
		return nil, errors.New("skiplist/DeleteRange: comparator is not set (== nil)")
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()

	// Walk the levels and nodes until we find the node at the lowest level (0) that the comparator returns false
	// E.g., if comparator is BuiltinLessThan, then we find the node at the lowest level s.t. node.key < key
	// Then we walk from there to find all the nodes that have node.key == key
	// We keep track of the last touched nodes at each level as selectFingers, and then we re-use the selectFingers
	// so that we can get O(log k) where k is the distance between last searched key and current search key
	// -- ok, so all this is done by updateSearchFingers

	if err = this.updateSearchFingers(key1, this.selectFingers); err != nil {
		return nil, errors.New("skiplist/DeleteRange: error finding node; " + err.Error())
	}

	iter = newIterator()
	var res bool
	for p := this.selectFingers[0].next[0]; p != nil; p = p.next[0] {
		pk := p.GetKey()
		if res, err = this.compare(pk, key2); err != nil {
			// If there's error in comparing the keys, then return err
			return nil, errors.New("skiplist/DeleteRange: error comparing keys; " + err.Error())
		} else if res || reflect.DeepEqual(pk, key2) {
			iter.buf = append(iter.buf, p)
			iter.count++

			for i := 0; i < this.level; i++ {
				if this.selectFingers[i].next[i] != p {
					break
				}
				this.selectFingers[i].next[i] = p.next[i]
			}

			this.count--

			for this.level > 1 && this.headNode.next[this.level-1] == nil {
				this.level--
			}
		} else {
			// Otherwise if the p.key is "after" key, after could mean greater or less, depending
			// on the comparator, then we know we are done
			break
		}
	}

	return iter, nil
}

func (this *Skiplist) RealCount(i int) (c int) {
	for p := this.headNode.next[i]; p != nil; {
		if p != nil {
			//log.Println("node =", p.record)
			c++
			p = p.next[i]
		}
	}

	return
}

func (this *Skiplist) PrintStats() {
	fmt.Println("Real count   :", this.RealCount(0))
	fmt.Println("Total levels :", this.Level())

	for i := 0; i < this.level; i++ {
		fmt.Println("Level", i, "count:", this.RealCount(i))
	}

}
