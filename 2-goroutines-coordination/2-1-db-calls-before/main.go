package main

import (
	"fmt"
	"time"
)

// slower
func main() {
	start := time.Now()

	user := getUser()
	orders := getOrders()
	payments := getPayments()

	fmt.Println(user)
	fmt.Println(orders)
	fmt.Println(payments)
	fmt.Println("Total time:", time.Since(start)) // ≈ 3 s
}

// mock “database” functions
func getUser() string {
	time.Sleep(1 * time.Second)
	return "User: Alice"
}

func getOrders() string {
	time.Sleep(1 * time.Second)
	return "Orders: [#123 #124]"
}

func getPayments() string {
	time.Sleep(1 * time.Second)
	return "Payments: Paid"
}
