/*
 * Copyright (c) 2013 Zhen, LLC. http://zhen.io. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license.
 *
 */

package skiplist

import (
	"reflect"
	"fmt"
)

var (
	BuiltinLessThan Comparator = func(k1, k2 interface{}) (bool, error) {
		t2 := reflect.TypeOf(k2).Name()

		switch k1 := k1.(type) {
		case string:
			if t2 != "string" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(string) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(string), nil

		case int64:
			if t2 != "int64" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(int64) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(int64), nil

		case int32:
			if t2 != "int32" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(int32) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(int32), nil

		case int16:
			if t2 != "int16" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(int16) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(int16), nil

		case int8:
			if t2 != "int8" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(int8) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(int8), nil

		case int:
			if t2 != "int" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(int) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(int), nil

		case float32:
			if t2 != "float32" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(float32) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(float32), nil

		case float64:
			if t2 != "float64" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(float64) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(float64), nil

		case uint:
			if t2 != "uint" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(uint) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(uint), nil

		case uint8:
			if t2 != "uint8" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(uint8) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(uint8), nil

		case uint16:
			if t2 != "uint16" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(uint16) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(uint16), nil

		case uint32:
			if t2 != "uint32" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(uint32) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(uint32), nil

		case uint64:
			if t2 != "uint64" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(uint64) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(uint64), nil

		case uintptr:
			if t2 != "uintptr" {
				return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: k1.(uintptr) and k2.(%s) have different types", t2)
			}
			return k1 < k2.(uintptr), nil
		}

		t1 := reflect.TypeOf(k1).Name()
		return false, fmt.Errorf("reducedb/skiplist:builtinLessThan: unsupported types for k1.(%s) and k2.(%s)", t1, t2)
	}

	BuiltinGreaterThan Comparator = func(k1, k2 interface{}) (bool, error) {
		t2 := reflect.TypeOf(k2).Name()

		switch k1 := k1.(type) {
		case string:
			if t2 != "string" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(string) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(string), nil

		case int64:
			if t2 != "int64" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(int64) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(int64), nil

		case int32:
			if t2 != "int32" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(int32) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(int32), nil

		case int16:
			if t2 != "int16" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(int16) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(int16), nil

		case int8:
			if t2 != "int8" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(int8) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(int8), nil

		case int:
			if t2 != "int" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(int) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(int), nil

		case float32:
			if t2 != "float32" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(float32) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(float32), nil

		case float64:
			if t2 != "float64" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(float64) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(float64), nil

		case uint:
			if t2 != "uint" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(uint) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(uint), nil

		case uint8:
			if t2 != "uint8" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(uint8) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(uint8), nil

		case uint16:
			if t2 != "uint16" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(uint16) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(uint16), nil

		case uint32:
			if t2 != "uint32" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(uint32) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(uint32), nil

		case uint64:
			if t2 != "uint64" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(uint64) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(uint64), nil

		case uintptr:
			if t2 != "uintptr" {
				return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: k1.(uintptr) and k2.(%s) have different types", t2)
			}
			return k1 > k2.(uintptr), nil
		}

		t1 := reflect.TypeOf(k1).Name()
		return false, fmt.Errorf("reducedb/skiplist:builtinGreaterThan: unsupported types for k1.(%s) and k2.(%s)", t1, t2)
	}	
	
	BuiltinEqual Comparator = func(k1, k2 interface{}) (bool, error) {
		t2 := reflect.TypeOf(k2).Name()

		switch k1 := k1.(type) {
		case string:
			if t2 != "string" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(string) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(string), nil

		case int64:
			if t2 != "int64" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(int64) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(int64), nil

		case int32:
			if t2 != "int32" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(int32) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(int32), nil

		case int16:
			if t2 != "int16" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(int16) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(int16), nil

		case int8:
			if t2 != "int8" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(int8) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(int8), nil

		case int:
			if t2 != "int" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(int) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(int), nil

		case float32:
			if t2 != "float32" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(float32) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(float32), nil

		case float64:
			if t2 != "float64" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(float64) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(float64), nil

		case uint:
			if t2 != "uint" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(uint) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(uint), nil

		case uint8:
			if t2 != "uint8" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(uint8) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(uint8), nil

		case uint16:
			if t2 != "uint16" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(uint16) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(uint16), nil

		case uint32:
			if t2 != "uint32" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(uint32) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(uint32), nil

		case uint64:
			if t2 != "uint64" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(uint64) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(uint64), nil

		case uintptr:
			if t2 != "uintptr" {
				return false, fmt.Errorf("reducedb/skiplist:builtinEqual: k1.(uintptr) and k2.(%s) have different types", t2)
			}
			return k1 == k2.(uintptr), nil
		}

		t1 := reflect.TypeOf(k1).Name()
		return false, fmt.Errorf("reducedb/skiplist:builtinEqual: unsupported types for k1.(%s) and k2.(%s)", t1, t2)
	}
)
