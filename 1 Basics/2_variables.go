package main

import "fmt"

func main() {
	// Declare
	var name string

	// Assign
	name = "Reassigned"

	// Automatic Type Inference
	var anotherName = "Kevin Reassigned"

	// Without var
	nameWithoutVar := "Kevin Declared without var"

	// Initialize Many
	var (
		one = "One"
		two = "Two"
	)

	// Constant
	const constantName = "Constant Kevin"

	fmt.Println(name)
	fmt.Println(anotherName)
	fmt.Println(nameWithoutVar)
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(constantName)
}