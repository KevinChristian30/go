package main

import "fmt"

// In Go, a closure is a function value that references variables from outside its body.
// These referenced variables are "closed over" by the function,
// meaning they remain accessible and can retain their values
// even after the function that created them has finished executing.

// Key Concepts of Closures in Go

// 1. Function as First-Class Citizen: Functions can be assigned to variables, passed as arguments, and returned from other functions.
// 2. Capturing Variables: A closure can capture and access variables defined outside its immediate scope.
// 3. Persistent State: The captured variables retain their values between calls to the closure.

func createCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	counter := createCounter()
	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
	fmt.Println(counter()) // Output: 3
}
