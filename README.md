Skiplist
========

Go implementation of skiplist, with search fingers.

Reference: http://drum.lib.umd.edu/bitstream/1903/544/2/CS-TR-2286.1.pdf 

This implementation supports duplicate keys. However, because of that, this implementation does not support Update.
### Examples

#### Woring with a skiplist of ints

```
// Creating a new skiplist, using the built-in Less Than function as the comparator.
// There are also two other built in comparators: BuiltinGreaterThan, BuiltinEqual
list := New(skiplist.BuiltinLessThan)

// Inserting key, value pairs into the skiplist. The skiplist is sorted by key,
// using the comparator function to determine order
list.Insert(1,1)
list.Insert(1,2)
list.Insert(2,1)
list.Insert(2,2)
list.Insert(2,3)
list.Insert(2,4)
list.Insert(2,5)
list.Insert(1,3)
list.Insert(1,4)
list.Insert(1,5)

// Selecting items that have the key == 1. Select returns a Skiplist.Iterator
rIter, err := list.Select(1)

// Iterate through the list of items. Keys and Values are turned as interface{}, so you
// need to type assert them to your type
for rIter.Next() {
	fmt.Println(rIter.Key().(int), rIter.Value().(int))
}

// Delete the items that match key. An iterator is returned with the list of deleted items.
rIter, err = list.Delete(1)

// You can also SelectRange or DeleteRange
rIter, err = list.SelectRange(1, 2)

rIter, err = list.DeleteRange(1, 2)
```

### Bultin Comparators

There are three built-in comparator functions:

* BuiltinLessThan: if you want to sort the skiplist in ascending order
* BuiltinGreaterThan: if you want to sort the skiplist in descending order
* BuiltinEqual: just to compare

Currently these built-in comparator functions work for all built-in Go types, including:

* string
* uint64, uint32, uint16, uint8, uint
* int64, int32, int16, int8, int
* float32, float64
* unitptr

### Performance

These numbers are from my Macbook Air 10.8.4 1.8GHz Intel Core i5 4GB 1600MHz DDR3

```
$ go test -v -bench=.
=== RUN TestInsertIntAscending
--- PASS: TestInsertIntAscending (0.32 seconds)
=== RUN TestInsertIntDescending
--- PASS: TestInsertIntDescending (0.35 seconds)
=== RUN TestInsertTimeAscending
--- PASS: TestInsertTimeAscending (0.20 seconds)
=== RUN TestInsertTimeDescending
--- PASS: TestInsertTimeDescending (0.11 seconds)
=== RUN TestInsertStringAscending
--- PASS: TestInsertStringAscending (0.47 seconds)
=== RUN TestInsertStringDescending
--- PASS: TestInsertStringDescending (0.48 seconds)
PASS
BenchmarkInsertTimeDescending    1000000              1252 ns/op
BenchmarkInsertTimeAscending     1000000              2211 ns/op
BenchmarkInsertInt               1000000              4722 ns/op
BenchmarkInsertInt64             1000000              4759 ns/op
BenchmarkInsertString            1000000              6995 ns/op
ok      github.com/zhenjl/skiplist      20.195s
```

Notice the fastest time is BenchmarkInsertTimeDescending. This is because the keys for that test is generated using time.Now().UnixNano(), which is always ascending. And because the sort order of the skiplist is descending, so the new item is ALWAYS inserted at the front of the list. This happens to have the best case of O(1).

The next best time is BenchmarkInsertTimeAscending. This is still pretty good, but because the sort order is ascending, so the new items are ALWAYS inserted at the end. This required the skiplist to walk all the levels so it took a bit longer.

The other benchmarks should have the average O(log k) efficiency.