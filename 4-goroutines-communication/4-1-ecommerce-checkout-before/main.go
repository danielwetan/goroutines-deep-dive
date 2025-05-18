package main

import (
	"fmt"
	"strconv"
	"time"
)

var productPrices = map[string]map[string]string{
	"storeA": {
		"orange": "10",
		"mango":  "20",
	},
	"storeB": {
		"laptop": "1000",
	},
}

// getPrice returns total price for a store's items
func getPrice(store string, items []string) int {
	fmt.Printf("[%s] Calculating price for items: %v\n", store, items)
	time.Sleep(1 * time.Second) // simulate delay

	total := 0
	storeItems := productPrices[store]
	for _, item := range items {
		if price, exists := storeItems[item]; exists {
			priceInt, _ := strconv.Atoi(price)
			total += priceInt
		}
	}
	return total
}

// sendInvoices loops through all store totals and simulates sending invoices
func sendInvoices(stores []string, totals map[string]int) {
	for _, store := range stores {
		total := totals[store]
		fmt.Printf("[%s] Sending invoice for $%d\n", store, total)
		time.Sleep(1 * time.Second) 
		fmt.Printf("[%s] Invoice sent!\n", store)
	}
}

func main() {
	start := time.Now()

	stores := []string{"storeA", "storeB"}
	totals := make(map[string]int)

	totals["storeA"] = getPrice("storeA", []string{"orange", "mango"})
	totals["storeB"] = getPrice("storeB", []string{"laptop"})

	sendInvoices(stores, totals)
	fmt.Println("Total time:", time.Since(start)) // ≈ 4s
}
