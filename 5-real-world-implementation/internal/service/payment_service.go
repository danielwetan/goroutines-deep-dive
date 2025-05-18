package service

import (
	"sync"

	"github.com/danielwetan/gin-clean-architecture/internal/model"
	"github.com/danielwetan/gin-clean-architecture/internal/repository"
)

type PaymentService interface {
	GetPaymentDetail() (*model.PaymentDetail, error)
	Checkout(req []model.CheckoutRequest) ([]model.CheckoutResponse, error)
}

type paymentService struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return &paymentService{paymentRepository: repo}
}

func (s *paymentService) GetPaymentDetail() (*model.PaymentDetail, error) {
	var wg sync.WaitGroup
	var product *model.Product
	var payment *model.Payment
	wg.Add(2)

	go func() {
		product, _ = s.paymentRepository.GetProduct()
		wg.Done()
	}()
	go func() {
		payment, _ = s.paymentRepository.GetPayment()
		wg.Done()
	}()

	wg.Wait()

	return &model.PaymentDetail{
		Product: *product,
		Payment: *payment,
	}, nil
}

func (s *paymentService) Checkout(req []model.CheckoutRequest) ([]model.CheckoutResponse, error) {
	invoiceChannel := make(chan [2]interface{})

	for _, data := range req {
		go s.paymentRepository.GetPrice(data.StoreName, data.Products, invoiceChannel)
	}

	// before
	// go s.paymentRepository.GetPrice("storeA", []string{"orange", "mango"}, invoiceChannel)
	// go s.paymentRepository.GetPrice("storeB", []string{"laptop"}, invoiceChannel)

	checkout := s.paymentRepository.SendInvoice(invoiceChannel)
	return checkout, nil
}
