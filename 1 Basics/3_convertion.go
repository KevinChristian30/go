package main

import "fmt"

func main() {
    var intVal int = 42
    var floatVal float64 = float64(intVal) // Cast int to float64

    fmt.Println("Integer:", intVal)  // Output: Integer: 42
    fmt.Println("Float:", floatVal)  // Output: Float: 42.0
}
