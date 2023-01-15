// Copyright 2013 The Go Authors. All rights reserved.
// Copyright 2022 Peter Aronoff. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package genericlist_test

import (
	"fmt"

	"github.com/telemachus/genericlist"
)

func Example() {
	// Create a new list and put some numbers in it.
	l := genericlist.New[int]()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents squared.
	squared := func(x int) int {
		return x * x
	}
	for e := l.Front(); e != nil; e = e.Next() {
		// NB: container/list treats all element values as interface{}.
		// In order to feed such a value to squared, the caller would
		// need to use a type assertion: squared(e.Value.(int)). In this
		// library, however, element values belong to a specific type.
		fmt.Printf("%d squared = %d\n", e.Value, squared(e.Value))
	}

	// Output:
	// 1 squared = 1
	// 2 squared = 4
	// 3 squared = 9
	// 4 squared = 16
}
