/*
 * Copyright (c) 2013 Dataence, LLC. http://zhen.io. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license.
 *
 */

package skiplist

type node struct {
	next  []*node
	key   interface{}
	value interface{}
}

// Create a new node with l levels of pointers
func newNode(l int) *node {
	return &node{
		next: make([]*node, l),
	}
}

func (this *node) SetKey(key interface{}) {
	this.key = key
}

func (this *node) GetKey() (key interface{}) {
	return this.key
}

func (this *node) SetValue(value interface{}) {
	this.value = value
}

func (this *node) GetValue() (key interface{}) {
	return this.value
}

func (this *node) Next() *node {
	return this.next[0]
}

func (this *node) NextAtLevel(l int) *node {
	if l >= 0 && l < len(this.next) {
		return this.next[l]
	}

	return nil
}
