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
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/dataence/compare"
)

func TestInsertIntAscending(t *testing.T) {
	count := 100000
	list := New(compare.BuiltinLessThan)
	keys := make([]int, count)

	for i := 0; i < count; i++ {
		keys[i] = rand.Intn(count)
	}

	for i := 0; i < count; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	//list.PrintStats()
	rc := list.RealCount(0)
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
	list := New(compare.BuiltinGreaterThan)
	keys := make([]int, count)

	for i := 0; i < count; i++ {
		keys[i] = rand.Intn(count)
	}

	for i := 0; i < count; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	//list.PrintStats()
	rc := list.RealCount(0)
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
	list := New(compare.BuiltinLessThan)
	keys := make([]int64, count)

	for i := 0; i < count; i++ {
		keys[i] = time.Now().UnixNano()
	}

	for i := 0; i < count; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	//list.PrintStats()
	rc := list.RealCount(0)
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
	list := New(compare.BuiltinGreaterThan)
	keys := make([]int64, count)

	for i := 0; i < count; i++ {
		keys[i] = time.Now().UnixNano()
	}

	for i := 0; i < count; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	//list.PrintStats()
	rc := list.RealCount(0)
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
	list := New(compare.BuiltinLessThan)
	keys := make([]string, count)

	for i := 0; i < count; i++ {
		j := rand.Intn(count)
		keys[i] = strconv.FormatInt(int64(j), 10)
	}

	for i := 0; i < count; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	//list.PrintStats()
	rc := list.RealCount(0)
	if rc != count {
		t.Fatal("Count not the same")
	}

	a := ""
	for p := list.headNode.next[0]; p != nil; p = p.next[0] {
		if greater, _ := compare.BuiltinGreaterThan(a, p.key); greater {
			t.Fatal(a, " >", p.key.(string))
		}
	}
}

func TestInsertStringDescending(t *testing.T) {
	count := 100000
	list := New(compare.BuiltinGreaterThan)
	keys := make([]string, count)

	for i := 0; i < count; i++ {
		j := rand.Intn(count)
		keys[i] = strconv.FormatInt(int64(j), 10)
	}

	for i := 0; i < count; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	//list.PrintStats()
	rc := list.RealCount(0)
	if rc != count {
		t.Fatal("Count not the same")
	}

	a := "zzzz"
	for p := list.headNode.next[0]; p != nil; p = p.next[0] {
		if less, _ := compare.BuiltinLessThan(a, p.key); less {
			t.Fatal(a, " <", p.key.(string))
		}
	}
}

func TestSelectInt(t *testing.T) {
	list := New(compare.BuiltinLessThan)

	list.Insert(1, 1)
	list.Insert(1, 2)
	list.Insert(2, 1)
	list.Insert(2, 2)
	list.Insert(2, 3)
	list.Insert(2, 4)
	list.Insert(2, 5)
	list.Insert(1, 3)
	list.Insert(1, 4)
	list.Insert(1, 5)

	rIter, _ := list.Select(1)

	if rIter.Count() != 5 {
		t.Fatal("number of results != 5")
	}

	//for rIter.Next() {
	//	fmt.Println(rIter.Key().(int), rIter.Value().(int))
	//}

	//rIter.Rewind()

	//for rIter.Next() {
	//	fmt.Println(rIter.Key().(int), rIter.Value().(int))
	//}
}

func TestSelectRangeInt(t *testing.T) {
	list := New(compare.BuiltinLessThan)

	list.Insert(1, 10)
	list.Insert(1, 20)
	list.Insert(2, 1)
	list.Insert(2, 2)
	list.Insert(2, 3)
	list.Insert(2, 4)
	list.Insert(2, 5)
	list.Insert(1, 30)
	list.Insert(1, 40)
	list.Insert(1, 50)
	list.Insert(3, 5)
	list.Insert(4, 5)
	list.Insert(5, 5)
	list.Insert(6, 5)

	rIter, _ := list.SelectRange(1, 2)

	//for i := 0; i < len(results); i++ {
	//	fmt.Println(results[i].(int))
	//}

	if rIter.Count() != 10 {
		t.Fatal("number of results != 10")
	}
	//for rIter.Next() {
	//	fmt.Println(rIter.Key().(int), rIter.Value().(int))
	//}
}

func TestSelectRangeInt2(t *testing.T) {
	count := 10000
	list := New(compare.BuiltinLessThan)
	keys := make([]int, count)
	total := 0

	for i := 0; i < count; i++ {
		keys[i] = rand.Intn(count)
		if keys[i] >= 100 && keys[i] <= 2000 {
			total++
		}
	}

	for i := 0; i < count; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	rIter, _ := list.SelectRange(100, 2000)

	//for i := 0; i < len(results); i++ {
	//	fmt.Println(results[i].(int))
	//}

	if rIter.Count() != total {
		t.Fatal("number of results !=", total)
	}
}

func TestDeleteInt(t *testing.T) {
	list := New(compare.BuiltinLessThan)

	list.Insert(1, 1)
	list.Insert(1, 2)
	list.Insert(2, 1)
	list.Insert(2, 2)
	list.Insert(2, 3)
	list.Insert(2, 4)
	list.Insert(2, 5)
	list.Insert(1, 3)
	list.Insert(1, 4)
	list.Insert(1, 5)

	fmt.Println("---")
	list.PrintStats()
	rIter, _ := list.Delete(1)

	if rIter.Count() != 5 {
		t.Fatal("number of results != 5")
	}

	dIter, _ := list.Select(1)

	if dIter.Count() != 0 {
		t.Fatal("still has key == 1")
	}
	fmt.Println("---")
	list.PrintStats()

}

func TestDeleteRangeInt2(t *testing.T) {
	count := 10000
	list := New(compare.BuiltinLessThan)
	keys := make([]int, count)
	total := 0

	for i := 0; i < count; i++ {
		keys[i] = rand.Intn(count)
		if keys[i] >= 100 && keys[i] <= 20000 {
			total++
		}
	}

	for i := 0; i < count; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			t.Fatal(err)
		}
	}

	fmt.Println("---")
	list.PrintStats()

	rIter, _ := list.DeleteRange(100, 20000)

	//for i := 0; i < len(results); i++ {
	//	fmt.Println(results[i].(int))
	//}

	if rIter.Count() != total {
		t.Fatal("number of results !=", total)
	}

	rIter, _ = list.SelectRange(100, 20000)
	if rIter.Count() != 0 {
		t.Fatal("still has nodes btwn 100 and 20000", rIter.Count())
	}

	if list.Count() != count-total {
		t.Fatal("remaining count !=", list.Count(), count-total)
	}

	fmt.Println("---")
	list.PrintStats()
}

func BenchmarkInsertTimeDescending(b *testing.B) {
	list := New(compare.BuiltinGreaterThan)
	keys := make([]int64, b.N)

	for i := 0; i < b.N; i++ {
		keys[i] = time.Now().UnixNano()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			b.Fatal(err)
		}
	}

	//for p := list.headNode.next[0]; p != nil; p = p.next[0] {
	//	fmt.Println(p.key.(int64))
	//}
}

func BenchmarkInsertTimeAscending(b *testing.B) {
	list := New(compare.BuiltinLessThan)
	keys := make([]int64, b.N)

	for i := 0; i < b.N; i++ {
		keys[i] = time.Now().UnixNano()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			b.Fatal(err)
		}
	}

	//for p := list.headNode.next[0]; p != nil; p = p.next[0] {
	//	fmt.Println(p.key.(int64))
	//}
}

func BenchmarkInsertInt(b *testing.B) {
	list := New(compare.BuiltinLessThan)
	keys := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = rand.Intn(b.N)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkInsertInt64(b *testing.B) {
	list := New(compare.BuiltinLessThan)
	keys := make([]int64, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = int64(rand.Intn(b.N))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkInsertString(b *testing.B) {
	list := New(compare.BuiltinLessThan)
	keys := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		j := rand.Intn(b.N)
		keys[i] = strconv.FormatInt(int64(j), 10)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := list.Insert(keys[i], i); err != nil {
			b.Fatal(err)
		}
	}
}
