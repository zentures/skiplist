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
