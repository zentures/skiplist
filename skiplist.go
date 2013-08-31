/*
 * Copyright (c) 2013 Zhen, LLC. http://zhen.io. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license.
 *
 */

package skiplist

import (
	"math/rand"
	"sync"
	//"log"
	"errors"
	"math"
	"fmt"
)

var (
	DefaultMaxLevel int = 12
	DefaultProbability float32 = 0.25
)

// Comparator for the node keys.
// For ascending order - if k1 < k2 return true; else return false
// For descending order - if k1 > k2 return true; else return false
type Comparator func(k1, k2 interface{}) (bool, error)

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
	fingers []*node

	// Total number of nodes inserted
	count int

	// Keep track of the number of nodes for each level
	levelCount []int

	// Comparison function
	compare Comparator


	mutex sync.RWMutex
}

func New(compare Comparator) *Skiplist {
	l := DefaultMaxLevel
	ip := int(math.Ceil(1/float64(DefaultProbability)))

	return &Skiplist{
		ip: ip,
		maxLevel: l,
		levelCount: make([]int, l),
		fingers: make([]*node, l),
		level: 1,
		count: 0,
		compare: compare,
		headNode: newNode(l),
	}
}

func (this *Skiplist) SetCompare(compare Comparator) (err error) {
	this.compare = compare
	return nil
}

func (this *Skiplist) SetMaxLevel(l int) (err error) {
	this.maxLevel = l
	return nil
}

func (this *Skiplist) SetProbability(p float32) (err error) {
	this.ip = int(math.Ceil(1/float64(p)))
	return nil
}

// Choose the new node's level, branching with p (1/ip) probability, with no regards to N (size of list)
func (this *Skiplist) newNodeLevel() int {
	h := 1

	for h < this.maxLevel && rand.Intn(this.ip) == 0 {
		h++
	}

	return h
}

func (this *Skiplist) updateSearchFingers(key interface{}) (err error) {
	startLevel := this.level-1
	startNode := this.headNode

	//log.Println("startLevel =", startLevel)
	//log.Println("startNode =", startNode)

	// Using Search Fingers
	// Reference: http://drum.lib.umd.edu/bitstream/1903/544/2/CS-TR-2286.1.pdf - section 3.1
	// if this.fingers[0] != nil && this.fingers[0].key < key {
	//log.Println("this.fingers[0] = ", this.fingers[0])


	if this.fingers[0] != nil && this.fingers[0].key != nil {
		if less, err := this.compare(this.fingers[0].key, key); err != nil {
			return err
		} else if less {
			// Move forward, find the highest level s.t. the next node's key < key
			for l := 1; l < this.level; l++ {
				if this.fingers[l].next[l] != nil && this.fingers[l].key != nil {
					// If the next node is not nil and this.fingers[l].key >= key
					if less, err := this.compare(this.fingers[l].key, key); err != nil {
						return err
					} else if less == false {
						startLevel = l - 1
						startNode = this.fingers[l]
						break
					}
				}
			}
		} else {
			//log.Println("inside if else, this.level =", this.level-1)
			// Move backward, find the lowest level s.t. the node's timestamp < t
			for l := 1; l < this.level; l++ {
				//log.Println("inside for loop, level =", l)
				// this.fingers[l].key < key
				if this.fingers[l].key != nil {
					if less, err := this.compare(this.fingers[l].key, key); err != nil {
						return err
					} else if less {
						startLevel = l
						startNode = this.fingers[l]
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
		//log.Println("inside for l,p loop, l =", l)
		//log.Println("inside for l,p loop, p =", p)
		n := p.next[l]
		//log.Println(n)

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

		this.fingers[l] = p
	}

	return nil
}

func (this *Skiplist) Insert(key, value interface{}) (err error) {

	// Create new node
	l := this.newNodeLevel()
	n := newNode(l)
	n.SetKey(key)
	n.SetValue(value)

	// If each Skiplist is single thread only, then we really don't need this..but for now let's keep
	this.mutex.Lock()
	defer this.mutex.Unlock()

	//log.Println("new node level =", l)
	this.levelCount[l-1]++

	//log.Println("this.finger[0] =", this.fingers[0])
	// Find the position where we should insert the node by updating the search fingers using the key
	// Search fingers will be updated with the rightmost element of each level that is left of the element
	// that's greater than or equal to key.
	// In other words, we are inserting the new node to the right of the search fingers.
	if err = this.updateSearchFingers(key); err != nil {
		return errors.New("reducedb/Skiplist:insert: cannot find insert position, " + err.Error())
	}

	//log.Println("search fingers =", this.fingers)
	// Raise the level of the skiplist if the new level is higher than the existing list level
	// So for levels higher than the current list level, the previous node is headNode for that level
	if this.level < l {
		for i := this.level; i < l; i++ {
			//log.Println("before ---- ", this.fingers)
			this.fingers[i] = this.headNode
			//log.Println("after  ---- ", this.fingers)
		}
		this.level = l
		//log.Println("new this.level =", l)
		//log.Println("this.fingers =", this.fingers)
	}

	// Finally insert the node into the skiplist
	for i := 0; i < l; i++ {
		// new node points forward to the previous node's next node
		// previous node's next node points to the new node
		n.next[i], this.fingers[i].next[i] = this.fingers[i].next[i], n
	}

	//log.Println("this.headNode.next =", this.headNode.next)
	//log.Println("            n.next =", n.next)

	// Adding to the count
	this.count++

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

func (this *Skiplist) RealCount() (c int) {
	for p := this.headNode.next[0]; p != nil; {
		if p != nil {
			//log.Println("node =", p.record)
			c++
			p = p.next[0]
		}
	}

	return
}

func (this *Skiplist) PrintInfo() {
	fmt.Println("Total nodes  :", this.Count())
	fmt.Println("Real count   :", this.RealCount())
	fmt.Println("Total levels :", this.Level())

	for i := 0; i < this.level; i++ {
		fmt.Println("Level", i, "count:", this.levelCount[i])
	}

}
