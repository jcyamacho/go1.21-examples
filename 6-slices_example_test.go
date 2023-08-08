package main

import (
	"cmp"
	"fmt"
	"slices"
)

// The new slices package provides many common operations on slices, using generic functions that work with slices of any element type.
//
// https://tip.golang.org/doc/go1.21#slices
//
// for all methods check documentation: https://pkg.go.dev/slices

func Example_slices_Contains() {
	s := []int{1, 2, 3, 4, 5}

	// Contains reports whether v is present in s.
	a := slices.Contains(s, 4)

	fmt.Println(a)

	// Output:
	// true
}

func Example_slices_ContainsFunc() {
	numbers := []int{0, 42, -10, 8}

	hasNegative := slices.ContainsFunc(numbers, func(n int) bool {
		return n < 0
	})

	fmt.Println(hasNegative)

	// Output:
	// true
}

func Example_slices_Min() {
	s := []int{6, 4, 7, 3, 8, 9}

	// Min returns the smallest element of s.
	a := slices.Min(s)

	fmt.Println(a)

	// Output:
	// 3
}

func Example_slices_Reverse() {
	names := []string{"alice", "Bob", "Vera"}

	slices.Reverse(names)

	fmt.Println(names)

	// Output:
	// [Vera Bob alice]
}

func Example_slices_BinarySearch() {
	names := []string{"Alice", "Bob", "Vera"}

	// BinarySearch searches for target in a sorted slice and returns the position where target is found,
	// or the position where target would appear in the sort order; it also returns a bool saying whether the target is really found in the slice.
	// The slice must be sorted in increasing order.

	n, found := slices.BinarySearch(names, "Vera")
	fmt.Println("Vera:", n, found)

	n, found = slices.BinarySearch(names, "Bill")
	fmt.Println("Bill:", n, found)

	// Output:
	// Vera: 2 true
	// Bill: 0 false
}

func Example_slices_BinarySearchFunc() {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 55},
		{"Bob", 24},
		{"Gopher", 13},
	}

	// BinarySearchFunc works like BinarySearch, but uses a custom comparison function.
	n, found := slices.BinarySearchFunc(people, Person{"Bob", 0}, func(a, b Person) int {
		return cmp.Compare(a.Name, b.Name)
	})

	fmt.Println("Bob:", n, found)

	// Output:
	// Bob: 1 true
}

func Example_slices_Clip() {
	s := make([]int, 5, 10)

	fmt.Println("len:", len(s), "cap:", cap(s))

	// Clip removes unused capacity from the slice.
	s = slices.Clip(s)

	fmt.Println("len:", len(s), "cap:", cap(s))

	// Output:
	// len: 5 cap: 10
	// len: 5 cap: 5
}

func Example_slices_Clone() {
	s1 := []int{1, 2, 3}

	// Clone returns a copy of the slice. The elements are copied using assignment, so this is a shallow clone.
	s2 := slices.Clone(s1)

	fmt.Println(s2)

	// Output:
	// [1 2 3]
}

func Example_slices_Compact() {
	s := []int{1, 1, 1, 2, 0, 3, 3, 0, 0, 0}

	// Compact replaces consecutive runs of equal elements with a single copy
	s = slices.Compact(s)

	fmt.Println(s)

	// Output:
	// [1 2 0 3 0]
}

func Example_slices_Compare() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{1, 2, 4}

	// Compare compares the elements of s1 and s2, using cmp.Compare on each pair of elements.
	// The elements are compared sequentially, starting at index 0, until one element is not equal to the other.
	// The result of comparing the first non-matching elements is returned.

	cm1 := slices.Compare(s1, s2)
	fmt.Println(cm1)

	cm2 := slices.Compare(s1, s3)
	fmt.Println(cm2)

	// Output:
	// 0
	// -1
}

func Example_slices_Delete() {
	s := []string{"a", "b", "c", "d", "e"}

	// Delete removes the elements s[i:j] from s, returning the modified slice
	s = slices.Delete(s, 1, 4)

	fmt.Println(s)

	// Output:
	// [a e]
}
