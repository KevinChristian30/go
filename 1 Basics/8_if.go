package main

import "fmt"

func main() {
	name := "Not Kevin a"

	if name == "Kevin" {
		fmt.Println(name)
	} else if name == "Not Kevin" {
		fmt.Println("Not Kevin")
	} else {
		fmt.Println("Fallback")
	}
}