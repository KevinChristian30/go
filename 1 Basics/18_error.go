package main

import (
	"fmt"
)

// Custom error
type DivideByZeroError struct {
}

// Custom error is created by implementing the Error interface
func (error *DivideByZeroError) Error() string {
	return "Can't Divide by Zero"
}

func Divide(numerator int, denumerator int) (int, error) {
	if denumerator == 0 {
		return -1, new(DivideByZeroError)
	}

	return numerator / denumerator, nil
}

func main() {
	value, err := Divide(1, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
