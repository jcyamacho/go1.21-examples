package main

import (
	"fmt"
)

func Example_max_int() {
	a := max(6, 7, 3, 5, 9)

	fmt.Println(a)

	// Output:
	// 9
}

func Example_max_float() {
	a := max(6.1, 7.2)

	fmt.Println(a)

	// Output:
	// 7.2
}

func Example_max_string() {
	a := max("foo", "bar")

	fmt.Println(a)

	// Output:
	// foo
}
