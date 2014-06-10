/*
 * Copyright (c) 2013 Dataence, LLC. http://zhen.io. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license.
 *
 */

package skiplist

import (
	"fmt"
	"reflect"
)

var (
	BuiltinLessThan    Comparator = builtinLessThan
	BuiltinGreaterThan Comparator = builtinGreaterThan
	BuiltinEqual       Comparator = builtinEqual
)

func builtinLessThan(k1, k2 interface{}) (bool, error) {
	if reflect.TypeOf(k1) != reflect.TypeOf(k2) {
		return false, fmt.Errorf("skiplist/BuiltinLessThan: k1.(%s) and k2.(%s) have different types",
			reflect.TypeOf(k1).Name(), reflect.TypeOf(k2).Name())
	}

	switch k1 := k1.(type) {
	case string:
		return k1 < k2.(string), nil

	case int64:
		return k1 < k2.(int64), nil

	case int32:
		return k1 < k2.(int32), nil

	case int16:
		return k1 < k2.(int16), nil

	case int8:
		return k1 < k2.(int8), nil

	case int:
		return k1 < k2.(int), nil

	case float32:
		return k1 < k2.(float32), nil

	case float64:
		return k1 < k2.(float64), nil

	case uint:
		return k1 < k2.(uint), nil

	case uint8:
		return k1 < k2.(uint8), nil

	case uint16:
		return k1 < k2.(uint16), nil

	case uint32:
		return k1 < k2.(uint32), nil

	case uint64:
		return k1 < k2.(uint64), nil

	case uintptr:
		return k1 < k2.(uintptr), nil
	}

	return false, fmt.Errorf("skiplist/BuiltinLessThan: unsupported types for k1.(%s) and k2.(%s)",
		reflect.TypeOf(k1).Name(), reflect.TypeOf(k2).Name())
}

func builtinGreaterThan(k1, k2 interface{}) (bool, error) {
	if reflect.TypeOf(k1) != reflect.TypeOf(k2) {
		return false, fmt.Errorf("skiplist/BuiltinGreaterThan: k1.(%s) and k2.(%s) have different types",
			reflect.TypeOf(k1).Name(), reflect.TypeOf(k2).Name())
	}

	switch k1 := k1.(type) {
	case string:
		return k1 > k2.(string), nil

	case int64:
		return k1 > k2.(int64), nil

	case int32:
		return k1 > k2.(int32), nil

	case int16:
		return k1 > k2.(int16), nil

	case int8:
		return k1 > k2.(int8), nil

	case int:
		return k1 > k2.(int), nil

	case float32:
		return k1 > k2.(float32), nil

	case float64:
		return k1 > k2.(float64), nil

	case uint:
		return k1 > k2.(uint), nil

	case uint8:
		return k1 > k2.(uint8), nil

	case uint16:
		return k1 > k2.(uint16), nil

	case uint32:
		return k1 > k2.(uint32), nil

	case uint64:
		return k1 > k2.(uint64), nil

	case uintptr:
		return k1 > k2.(uintptr), nil
	}

	return false, fmt.Errorf("skiplist/BuiltinGreaterThan: unsupported types for k1.(%s) and k2.(%s)",
		reflect.TypeOf(k1).Name(), reflect.TypeOf(k2).Name())
}

func builtinEqual(k1, k2 interface{}) (bool, error) {
	if reflect.TypeOf(k1) != reflect.TypeOf(k2) {
		return false, fmt.Errorf("skiplist/BuiltinEqual: k1.(%s) and k2.(%s) have different types",
			reflect.TypeOf(k1).Name(), reflect.TypeOf(k2).Name())
	}

	switch k1 := k1.(type) {
	case string:
		return k1 == k2.(string), nil

	case int64:
		return k1 == k2.(int64), nil

	case int32:
		return k1 == k2.(int32), nil

	case int16:
		return k1 == k2.(int16), nil

	case int8:
		return k1 == k2.(int8), nil

	case int:
		return k1 == k2.(int), nil

	case float32:
		return k1 == k2.(float32), nil

	case float64:
		return k1 == k2.(float64), nil

	case uint:
		return k1 == k2.(uint), nil

	case uint8:
		return k1 == k2.(uint8), nil

	case uint16:
		return k1 == k2.(uint16), nil

	case uint32:
		return k1 == k2.(uint32), nil

	case uint64:
		return k1 == k2.(uint64), nil

	case uintptr:
		return k1 == k2.(uintptr), nil
	}

	return false, fmt.Errorf("skiplist/BuiltinLessThan: unsupported types for k1.(%s) and k2.(%s)",
		reflect.TypeOf(k1).Name(), reflect.TypeOf(k2).Name())
}
