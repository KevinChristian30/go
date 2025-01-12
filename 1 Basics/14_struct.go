package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

// Give method to struct
func (human Human) sayHello() {
	fmt.Println("Hello, i am ", human.Name)
}

func main() {
	kevin := Human{
		Name: "Kevin",
		Age:  10,
	}

	fmt.Println(kevin)
	fmt.Println(kevin.Age)
	kevin.sayHello()
}
