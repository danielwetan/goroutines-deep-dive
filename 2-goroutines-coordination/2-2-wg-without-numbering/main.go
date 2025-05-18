package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var (
		wg sync.WaitGroup
	)

	// Create a slice of functions to execute
	tasks := []func() string{
		getUser,
		getOrders,
		// getPayments,
	}

	// Add the number of tasks automatically
	wg.Add(len(tasks))

	// Results slice to store the outputs in order
	results := make([]string, len(tasks))

	// Launch goroutines for each task
	for i, task := range tasks {
		go func(index int, fn func() string) {
			results[index] = fn()
			wg.Done()
		}(i, task)
	}

	wg.Wait()

	// Print results in order
	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Println("Total time:", time.Since(start)) // ≈ 1 s
}

// mock “database” functions
func getUser() string {
	time.Sleep(1 * time.Second)
	return "User: Adam"
}

func getOrders() string {
	time.Sleep(1 * time.Second)
	return "Orders: [#123 #124]"
}

func getPayments() string {
	time.Sleep(1 * time.Second)
	return "Payments: Paid"
}
