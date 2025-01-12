package main

import "fmt"

func main() {
    var intVal int = 42
    var floatVal float64 = float64(intVal) // Cast int to float64

    fmt.Println("Integer:", intVal)  // Output: Integer: 42
    fmt.Println("Float:", floatVal)  // Output: Float: 42.0

    // Type assertion
    // Type assertion is used to convert an empty interface/any to a type
    var result any = "Kevin"
    var assertedResult = result.(string)
    fmt.Println(assertedResult)
}
