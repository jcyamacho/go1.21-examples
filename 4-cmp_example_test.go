package main

import (
	"cmp"
	"fmt"
)

// The new cmp package defines the type constraint Ordered and two new generic functions Less and Compare that are useful with ordered types.
//
// https://tip.golang.org/doc/go1.21#cmp

// cmp.Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
func Min[T cmp.Ordered](a T, s ...T) T {
	if len(s) == 0 {
		return a
	}

	m := a

	for _, v := range s {
		if v < m {
			m = v
		}
	}

	return m
}

func ExampleLess() {
	x := 5
	y := 8

	// Less reports whether x is less than y
	a := cmp.Less(x, y)

	fmt.Println(a)

	// Output:
	// true
}

func ExampleCompare() {
	x := 3
	y := 7

	// Compare returns an integer comparing two values.
	// The result will be:
	//    0 if x == y
	//   -1 if x < y
	//   +1 if x > y

	a := cmp.Compare(x, y)

	fmt.Println(a)

	// Output:
	// -1
}
