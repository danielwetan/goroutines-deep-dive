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

// getPrice calculates total for a store and sends it to invoiceChannel
func getPrice(store string, items []string, invoiceChannel chan<- [2]interface{}) {
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

	invoiceChannel <- [2]interface{}{store, total}
}

// sendInvoice receives store + total and simulates invoice sending
func sendInvoice(invoiceChannel <-chan [2]interface{}) {
	for i := 0; i < 2; i++ { // expecting 2 stores
		data := <-invoiceChannel
		store := data[0].(string)
		total := data[1].(int)

		fmt.Printf("[%s] Sending invoice for $%d\n", store, total)
		time.Sleep(1 * time.Second)
		fmt.Printf("[%s] Invoice sent!\n", store)
	}
}

func main() {
	start := time.Now()

	invoiceChannel := make(chan [2]interface{})

	go getPrice("storeA", []string{"orange", "mango"}, invoiceChannel)
	go getPrice("storeB", []string{"laptop"}, invoiceChannel)

	sendInvoice(invoiceChannel)
	fmt.Println("Total time:", time.Since(start)) // ≈ 3 s
}
