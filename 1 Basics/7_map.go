package main

import "fmt"

func main() {
	myMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	fmt.Println(myMap)
	fmt.Println(myMap["apple"])
	delete(myMap, "apple")
	fmt.Println(myMap["apple"])
}