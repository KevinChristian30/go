package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "0"
	names[1] = "1"
	names[2] = "2"

	fmt.Println(names)

	var values = [...]int{
		1, 2,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}