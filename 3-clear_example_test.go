package main

import (
	"examples/print"
	"fmt"
)

// The built-in function clear takes an argument of map, slice, or type parameter type, and deletes or zeroes out all elements.
//
// https://tip.golang.org/ref/spec#Clear

func Example_clear_map() {
	m := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	// deletes all entries, resulting in an	empty map
	clear(m)

	fmt.Println(m)

	// Output:
	// map[]
}

func Example_clear_slice() {
	s := []int{1, 2, 3}

	// sets all elements up to the length of "s" to the zero value of T
	clear(s)

	fmt.Println(s)

	// Output:
	// [0 0 0]
}

func Example_clear_type_parameter() {
	a := []struct {
		Name string
		Age  int
	}{
		{"foo", 1},
		{"bar", 2},
	}

	// sets all elements up to the length of "a" to the zero value of T
	clear(a)

	print.JSON(a)

	// Output:
	// [{"Name":"","Age":0},{"Name":"","Age":0}]
}

func Example_reuse_slice() {
	s := []int{1, 2, 3}

	fmt.Printf("len(s) = %d, cap(s) = %d\n", len(s), cap(s))

	a := s[:0]

	fmt.Printf("len(a) = %d, cap(a) = %d\n", len(a), cap(a))

	a = append(a, 9)

	fmt.Printf("a = %v\n", a)
	fmt.Printf("s = %v\n", s)

	// Output:
	// len(s) = 3, cap(s) = 3
	// len(a) = 0, cap(a) = 3
	// a = [9]
	// s = [9 2 3]
}
