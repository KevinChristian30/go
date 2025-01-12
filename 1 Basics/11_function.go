package main

import "fmt"

func minAndMax(array []int) (int, int) {
	if len(array) == 0 {
		return 0, 0
	}

	var (
		min = array[0]
		max = array[0]
	)
	for i := 0; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
		} else if array[i] > max {
			max = array[i]
		}
	}

	return min, max
}

func namedReturnValues() (firstName string, lastName string) {
	firstName = "First"

	return firstName, lastName // lastname is an empty string
}

// Variadic Functions
func sumAll(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

// Function with Callback
func functionalProgrammingMap(array []int, callback func(value int) int) []int {
	var result []int
	for i := 0; i < len(array); i++ {
		result = append(result, callback(array[i]))
	}
	return result
}

func main() {
	array := []int{
		1, 2, 3, 4, 5,
	}

	min, max := minAndMax(array)
	fmt.Println(min, max)

	firstName, lastName := namedReturnValues()
	fmt.Println(firstName, lastName)

	fmt.Println(sumAll(1, 2, 3, 4, 5))

	// Storing function in variable
	minAndMax2 := minAndMax
	fmt.Println(minAndMax2(array))

	// Call function with callback
	fmt.Println(functionalProgrammingMap(array, func(value int) int {
		return value * 2
	}))
}
