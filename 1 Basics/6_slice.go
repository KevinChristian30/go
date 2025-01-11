package main

import "fmt"

func main() {
	// In Go, a slice is a flexible, dynamically-sized view of an array.
	// Unlike arrays, slices are more powerful and commonly used
	// because they allow for resizing and provide various utilities for handling collections of data.
	// Remember, slices are just pointers to arrays

	// Basics
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4] // Elements from index 1 to 3
	arr[2] = 1000
	fmt.Println(slice) // Output: [20 1000 40]

	// With make function
	sliceFromMake := make([]int, 3, 5) // Length 3, Capacity 5
	fmt.Println(sliceFromMake)         // Output: [0 0 0]
	fmt.Println(len(sliceFromMake))    // Output: 3
	fmt.Println(cap(sliceFromMake))    // Output: 5

	// Directly Creating Slice
	directlyCreatedSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(directlyCreatedSlice) // Output: [1 2 3 4 5]

	// Appending a slice
	appendedSlice := make([]int, 2, 3)
	fmt.Println(len(appendedSlice), cap(appendedSlice)) // Output: 2 3

	appendedSlice = append(appendedSlice, 1, 2)         // Triggers reallocation
	fmt.Println(len(appendedSlice), cap(appendedSlice)) // Output: 4 6 (capacity doubled)

	// Hard Copying Slice
	originalSlice := []int{1, 2, 3, 4, 5}
	hardCopiedSlice := make([]int, len(originalSlice))

	copy(hardCopiedSlice, originalSlice)

	// Modify the copy to demonstrate it's a hard copy
	hardCopiedSlice[0] = 10

	// Print both slices
	fmt.Println("Original slice:", originalSlice) // Output: [1 2 3 4 5]
	fmt.Println("Copy slice:", hardCopiedSlice)   // Output: [10 2 3 4 5]
}
