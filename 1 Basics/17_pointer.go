package main

import "fmt"

type Person struct {
	Name string
}

// It is recommended to use pointer to reduce memory usage
// By default is it pass by value, so a local variable is created
func (person *Person) UpdateName() {
	person.Name = "Changed"
}

func exampleOne() {
	person1 := Person{Name: "Kevin"}
	person2 := person1 // This creates a hard copy

	person2.Name = "Changed" // This doesn't change person1's name

	fmt.Println(person1)
	fmt.Println(person2)
}

func exampleTwo() {
	person1 := Person{Name: "Kevin"}
	var person2 *Person = &person1 // This is a reference

	person2.Name = "Changed" // This doesn't change person1's name

	fmt.Println(person1)
	fmt.Println(*person2)
}

func newKeyword() {
	var person1 *Person = new(Person)
	fmt.Println(person1) // Empty String
}

func method() {
	person := Person{Name: "Kevin"}
	person.UpdateName()
	fmt.Println(person)
}

func main() {
	// exampleOne()
	// exampleTwo()
	// newKeyword()
	method()
}
