package repository

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/danielwetan/gin-clean-architecture/internal/model"
)

type PaymentRepository interface {
	GetProduct() (*model.Product, error)
	GetPayment() (*model.Payment, error)
	GetPrice(store string, items []string, invoiceChannel chan<- [2]interface{})
	SendInvoice(invoiceChannel <-chan [2]interface{}) []model.CheckoutResponse
	GenerateInvoiceNumber() string
}

type inMemoryPaymentRepository struct {
	payments        map[string]*model.PaymentDetail
	mu              sync.RWMutex
	lastOrderNumber int
}

func NewInMemoryPaymentRepository() PaymentRepository {
	return &inMemoryPaymentRepository{
		payments: make(map[string]*model.PaymentDetail),
	}
}

func (r *inMemoryPaymentRepository) GetProduct() (*model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	// time.Sleep(200 * time.Millisecond)
	time.Sleep(1 * time.Second)
	product := &model.Product{
		Name:  "Product 1",
		Price: 100,
	}
	return product, nil
}

func (r *inMemoryPaymentRepository) GetPayment() (*model.Payment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	// time.Sleep(200 * time.Millisecond)
	time.Sleep(1 * time.Second)
	payment := &model.Payment{
		Status: "Paid",
	}
	return payment, nil
}

var productPrices = map[string]map[string]string{
	"storeA": {
		"orange": "10",
		"mango":  "20",
	},
	"storeB": {
		"laptop": "1000",
	},
	"mall": {
		"baju": "300",
	},
}

func (r *inMemoryPaymentRepository) GetPrice(store string, items []string, invoiceChannel chan<- [2]interface{}) {
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

func (r *inMemoryPaymentRepository) SendInvoice(invoiceChannel <-chan [2]interface{}) []model.CheckoutResponse {
	var response []model.CheckoutResponse

	for i := 0; i < 3; i++ {
		data := <-invoiceChannel
		store := data[0].(string) // store name
		total := data[1].(int)    // total

		response = append(response, model.CheckoutResponse{
			StoreName: store,
			InvoiceID: r.GenerateInvoiceNumber(),
		})

		fmt.Printf("[%s] Sending invoice for $%d\n", store, total)
		// kirim email
		// mailgun.SendEmail()

		time.Sleep(1 * time.Second)
		fmt.Printf("[%s] Invoice sent!\n", store)
	}

	return response
}

func (r *inMemoryPaymentRepository) GenerateInvoiceNumber() string {
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	r.lastOrderNumber++
	date := time.Now().Format("2006/01/02") // format: 2024/05/18
	return fmt.Sprintf("INV/%s/%04d", date, r.lastOrderNumber)
}
