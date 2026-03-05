package main

import (
	"fmt"
	"time"
)

func get() int {
	time.Sleep(3 * time.Second)
	return 1
}

func main() {
	// Create channel
	channel := make(chan int)
	defer close(channel)

	go func() {
		channel <- get()
	}()

	// Receive data, blocks until data is received
	fmt.Println("First")
	data := <-channel
	fmt.Println("Second")

	fmt.Println(data)
}
