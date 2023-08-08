package main

import (
	"fmt"
	"maps"
	"strings"
)

// The new maps package provides several common operations on maps, using generic functions that work with maps of any key or element type.
//
// https://tip.golang.org/doc/go1.21#maps

func Example_maps_Clone() {
	m := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	// Clone returns a copy of m. This is a shallow clone: the new keys and values are set using ordinary assignment.
	c := maps.Clone(m)

	fmt.Println(c)

	// Output:
	// map[bar:2 baz:3 foo:1]
}

func Example_maps_Copy() {
	dst := map[string]int{"foo": 1}
	src := map[string]int{"bar": 2, "baz": 3}

	// Copy copies all key/value pairs from src into dst. If a key is common to both maps, the value from src is used.
	maps.Copy(dst, src)

	fmt.Println(dst)

	// Output:
	// map[bar:2 baz:3 foo:1]
}

func Example_maps_DeleteFunc() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	// DeleteFunc deletes any key/value pairs from m for which del returns true.
	maps.DeleteFunc(m, func(k string, v int) bool {
		return v%2 != 0 // delete odd values
	})

	fmt.Println(m)

	// Output:
	// map[four:4 two:2]
}

func Example_maps_Equal() {
	m1 := map[string]int{"foo": 1, "bar": 2, "baz": 3}
	m2 := map[string]int{"foo": 1, "bar": 2, "baz": 3}

	// Equal reports whether two maps contain the same key/value pairs. Values are compared using ==
	eq := maps.Equal(m1, m2)

	fmt.Println(eq)

	// Output:
	// true
}

func Example_maps_EqualFunc() {
	m1 := map[int]string{
		1:    "one",
		10:   "Ten",
		1000: "THOUSAND",
	}
	m2 := map[int][]byte{
		1:    []byte("One"),
		10:   []byte("Ten"),
		1000: []byte("Thousand"),
	}

	// EqualFunc is like Equal, but compares values using eq. Keys are still compared with ==
	eq := maps.EqualFunc(m1, m2, func(v1 string, v2 []byte) bool {
		return strings.ToLower(v1) == strings.ToLower(string(v2))
	})

	fmt.Println(eq)

	// Output:
	// true
}
