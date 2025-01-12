package main

import "fmt"

type IHasName interface {
	GetName() string
}

func sayHello(hasName IHasName) {
	fmt.Println("Hello ", hasName.GetName())
}

// Implementation
// The implementation isn't explicit
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// Empty interface
// Go isn't object oriented
// In OOP languages, we usually have a parent of all objects, like java.lang.Object in java
// To handle this, Go has an empty interface, which doesn't have method declaration, which means any data type is it's implementation
// By default, empty interface has a type alias called "any"
// Example: fmt.Println(value ...any) or fmt.Println(value ...interface{})

func main() {
	person := Person{
		Name: "Kevin",
	}
	sayHello(person)
}
