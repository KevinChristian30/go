package main

import "fmt"

// Defer delays the execution of a function
func deferedFunction() {
	fmt.Println("Start")

	defer fmt.Println("Deferred 1") // Scheduled to run later
	defer fmt.Println("Deferred 2") // Scheduled after Deferred 1

	fmt.Println("End")
}

// Panic terminates the whole program, defer will still be ran
func panicedFunction() {
	fmt.Println("Before panic")

	panic("Something went wrong!") // Trigger a panic

	fmt.Println("This will not be printed") // Not run
}

func main() {
	// Handle panic with recover, because defer will still be ran, it is used to handle the panic
	// It must be registered first
	defer func() {
		if errorMessage := recover(); errorMessage != nil {
			fmt.Println(errorMessage)
		}
	}()

	deferedFunction()
	panicedFunction()
}
