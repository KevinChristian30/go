package main

import "fmt"

func printType(i interface{}) {
    switch v := i.(type) {
		case int:
			fmt.Println("Integer:", v)
		case string:
			fmt.Println("String:", v)
		case bool:
			fmt.Println("Boolean:", v)
		default:
			fmt.Println("Unknown type")
    }
}

func main() {
    printType(42)      // Output: Integer: 42
    printType("hello") // Output: String: hello
    printType(true)    // Output: Boolean: true
}
