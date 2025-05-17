package main

import (
	"fmt"
	"time"
)

var count int

// race condition
// not safe
func increment() {
	for i := 0; i < 1000; i++ {
		count++
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go increment()
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Final count:", count) // wrong!
	fmt.Println("Run detector with `go run -race`")
}
