/*
 * Copyright (c) 2013 Dataence, LLC. http://zhen.io. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license.
 *
 */

package skiplist

type Iterator struct {
	// buffered nodes
	buf []*node

	// total count
	count int

	// current position
	cur int
}

func newIterator() *Iterator {
	return &Iterator{
		buf:   make([]*node, 0, 50),
		count: 0,
		cur:   -1,
	}
}

func (this *Iterator) Next() bool {
	this.cur++
	if this.cur >= this.count {
		return false
	}
	return true
}

func (this *Iterator) Key() interface{} {
	if this.cur < 0 || this.cur >= this.count {
		return nil
	}
	return this.buf[this.cur].GetKey()
}

func (this *Iterator) Value() interface{} {
	if this.cur < 0 || this.cur >= this.count {
		return nil
	}
	return this.buf[this.cur].GetValue()
}

func (this *Iterator) Rewind() {
	this.cur = -1
}

func (this *Iterator) Count() int {
	return this.count
}
