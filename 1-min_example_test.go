package main

import (
	"fmt"
)

// The built-in functions min and max compute the smallest (or largest, respectively) value of a fixed number of arguments of ordered types.
// There must be at least one argument.
//
// https://tip.golang.org/ref/spec#Min_and_max

func Example_min_int() {
	a := min(6, 7, 3, 5, 9)

	fmt.Println(a)

	// Output:
	// 3
}

func Example_min_float() {
	a := min(6.1, 7.2)

	fmt.Println(a)

	// Output:
	// 6.1
}

func Example_min_string() {
	a := min("foo", "bar")

	fmt.Println(a)

	// Output:
	// bar
}
