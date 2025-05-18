package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lastOrderNumber int
	mu              sync.Mutex
)

func getNextOrderID() string {
	mu.Lock()
	defer mu.Unlock()

	lastOrderNumber++
	date := time.Now().Format("2006/01/02") // format: 2024/05/18
	return fmt.Sprintf("INV/%s/%04d", date, lastOrderNumber)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			orderID := getNextOrderID()
			fmt.Printf("Worker %d generated Order ID: %s\n", workerID, orderID)
		}(i)
	}

	wg.Wait()
}
