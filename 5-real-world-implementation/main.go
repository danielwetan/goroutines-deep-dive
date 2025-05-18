package main

import (
	"github.com/danielwetan/gin-clean-architecture/internal/controller"
	"github.com/danielwetan/gin-clean-architecture/internal/repository"
	"github.com/danielwetan/gin-clean-architecture/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize dependencies
	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Setup router
	r := gin.Default()

	// Routes
	r.POST("/users", userController.CreateUser)
	r.GET("/users/:id", userController.GetUser) // path parameter
	r.GET("/users", userController.ListUsers)

	// Payment Detail
	paymentRepo := repository.NewInMemoryPaymentRepository()
	paymentService := service.NewPaymentService(paymentRepo)
	paymentController := controller.NewPaymentController(paymentService)
	r.GET("/payment", paymentController.GetPaymentDetail)
	r.POST("/checkout", paymentController.Checkout)

	// Start server
	r.Run(":9000")
}
