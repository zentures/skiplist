/*
 * Copyright (c) 2013 Zhen, LLC. http://zhen.io. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license.
 *
 */

package skiplist

import (
	"testing"
	//"fmt"
	"strconv"
	"math/rand"
	"time"
)

func TestInsertIntAscending(t *testing.T) {
	count := 100000
	list := New(builtinLessThan)
	keys := make([]int, count)

	for i := 0; i < count; i++ {
		keys[i] = rand.Intn(count)
	}

	for i := 0; i < count; i++ {
		if err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	//list.PrintInfo()
	rc := list.RealCount()
	if rc != count {
		t.Fatal("Count not the same")
	}

	i := 0
	j := -1
	for p := list.headNode.next[0]; p != nil; p = p.next[0] {
		k := p.key.(int)
		if j > k {
			t.Fatal(j, " >", k)
		}
		i++
	}
}

func TestInsertIntDescending(t *testing.T) {
	count := 100000
	list := New(builtinGreaterThan)
	keys := make([]int, count)

	for i := 0; i < count; i++ {
		keys[i] = rand.Intn(count)
	}

	for i := 0; i < count; i++ {
		if err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	//list.PrintInfo()
	rc := list.RealCount()
	if rc != count {
		t.Fatal("Count not the same")
	}

	i := 0
	j := 100001
	for p := list.headNode.next[0]; p != nil; p = p.next[0] {
		k := p.key.(int)
		if j < k {
			t.Fatal(j, " <", k)
		}
		i++
	}
}

func TestInsertTimeAscending(t *testing.T) {
	count := 100000
	list := New(builtinLessThan)
	keys := make([]int64, count)

	for i := 0; i < count; i++ {
		keys[i] = time.Now().UnixNano()
	}

	for i := 0; i < count; i++ {
		if err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	//list.PrintInfo()
	rc := list.RealCount()
	if rc != count {
		t.Fatal("Count not the same")
	}

	i := int64(0)
	j := int64(-1)
	for p := list.headNode.next[0]; p != nil; p = p.next[0] {
		k := p.key.(int64)
		if j > k {
			t.Fatal(j, " >", k)
		}
		i++
	}
}

func TestInsertTimeDescending(t *testing.T) {
	count := 100000
	list := New(builtinGreaterThan)
	keys := make([]int64, count)

	for i := 0; i < count; i++ {
		keys[i] = time.Now().UnixNano()
	}

	for i := 0; i < count; i++ {
		if err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	//list.PrintInfo()
	rc := list.RealCount()
	if rc != count {
		t.Fatal("Count not the same")
	}

	i := int64(0)
	j := time.Now().UnixNano()
	for p := list.headNode.next[0]; p != nil; p = p.next[0] {
		k := p.key.(int64)
		if j < k {
			t.Fatal(j, " <", k)
		}
		i++
	}
}

func TestInsertStringAscending(t *testing.T) {
	count := 100000
	list := New(builtinLessThan)
	keys := make([]string, count)

	for i := 0; i < count; i++ {
		j := rand.Intn(count)
		keys[i] = strconv.FormatInt(int64(j), 10)
	}

	for i := 0; i < count; i++ {
		if err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	//list.PrintInfo()
	rc := list.RealCount()
	if rc != count {
		t.Fatal("Count not the same")
	}

	a := ""
	for p := list.headNode.next[0]; p != nil; p = p.next[0] {
		if greater, _ := builtinGreaterThan(a, p.key); greater {
			t.Fatal(a, " >", p.key.(string))
		}
	}
}

func TestInsertStringDescending(t *testing.T) {
	count := 100000
	list := New(builtinGreaterThan)
	keys := make([]string, count)

	for i := 0; i < count; i++ {
		j := rand.Intn(count)
		keys[i] = strconv.FormatInt(int64(j), 10)
	}

	for i := 0; i < count; i++ {
		if err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	//list.PrintInfo()
	rc := list.RealCount()
	if rc != count {
		t.Fatal("Count not the same")
	}

	a := "zzzz"
	for p := list.headNode.next[0]; p != nil; p = p.next[0] {
		if less, _ := builtinLessThan(a, p.key); less {
			t.Fatal(a, " <", p.key.(string))
		}
	}
}

func BenchmarkInsertInt(b *testing.B) {
	list := New(builtinLessThan)
	keys := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Intn(b.N)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := list.Insert(keys[i], i); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkInsertInt64(b *testing.B) {
	list := New(builtinLessThan)
	keys := make([]int64, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = int64(rand.Intn(b.N))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := list.Insert(keys[i], i); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkInsertString(b *testing.B) {
	list := New(builtinLessThan)
	keys := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		j := rand.Intn(b.N)
		keys[i] = strconv.FormatInt(int64(j), 10)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := list.Insert(keys[i], i); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkInsertTimeDescending(b *testing.B) {
	list := New(builtinGreaterThan)
	keys := make([]int64, b.N)

	for i := 0; i < b.N; i++ {
		keys[i] = time.Now().UnixNano()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := list.Insert(keys[i], i); err != nil {
			b.Fatal(err)
		}
	}

	//for p := list.headNode.next[0]; p != nil; p = p.next[0] {
	//	fmt.Println(p.key.(int64))
	//}
}
