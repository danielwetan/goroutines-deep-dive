package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	getUser()
	getOrders()
	fmt.Println("Main done (might finish first)")
	fmt.Println("Total time:", time.Since(start))
}

// mock db calls
func getUser() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("user: Adam")
}

func getOrders() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("orders: #123")
}
