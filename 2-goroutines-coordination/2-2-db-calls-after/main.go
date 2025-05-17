package main

import (
	"fmt"
	"sync"
	"time"
)

// faster
func main() {
	start := time.Now()

	var (
		user, orders, payments string
		wg                     sync.WaitGroup
	)
	wg.Add(3)

	go func() {
		user = getUser()
		wg.Done()
	}()
	go func() {
		orders = getOrders()
		wg.Done()
	}()
	go func() {
		payments = getPayments()
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(user)
	fmt.Println(orders)
	fmt.Println(payments)
	fmt.Println("Total time:", time.Since(start)) // ≈ 1 s
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
