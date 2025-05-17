package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	mu    sync.Mutex
)

// safe with mutex
func increment() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go increment()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Final count:", count) // always 10 000
}
