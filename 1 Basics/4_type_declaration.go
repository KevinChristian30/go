package main

import "fmt"

func main() {
	type ID string

	var id ID = "123"

	fmt.Println(id)
}