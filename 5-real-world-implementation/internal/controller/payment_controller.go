package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/danielwetan/gin-clean-architecture/internal/model"
	"github.com/danielwetan/gin-clean-architecture/internal/service"
	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentService service.PaymentService
}

func NewPaymentController(service service.PaymentService) *PaymentController {
	return &PaymentController{paymentService: service}
}

func (c *PaymentController) GetPaymentDetail(ctx *gin.Context) {
	payment, err := c.paymentService.GetPaymentDetail()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, payment)
}

func (c *PaymentController) Checkout(ctx *gin.Context) {
	start := time.Now()

	var checkoutRequest []model.CheckoutRequest
	if err := ctx.ShouldBindJSON(&checkoutRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkout, _ := c.paymentService.Checkout(checkoutRequest)

	fmt.Println("Total time:", time.Since(start))

	ctx.JSON(200, checkout)
}
