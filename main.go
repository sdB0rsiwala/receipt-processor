package main

import (
	"fmt"
	controller "receipt-processor/controllers"
	"receipt-processor/repository"
	"receipt-processor/routes"
	"receipt-processor/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// main.go
// This file initializes and starts the Receipt Processor API server.
// It sets up dependencies, configures routes, and starts the Gin HTTP server.

func main() {
	// Configure Logrus for structured logging.
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Starting Receipt Processor API....")

	// Initialize the repository, service, and controller layers.
	receiptRepo := repository.NewReceiptRepository()                     // In-memory storage for receipts
	receiptService := services.NewReceiptService(receiptRepo)            // Business logic for receipts
	receiptController := controller.NewReceiptController(receiptService) // API controller

	// Initialize Gin router for handling HTTP requests.
	router := gin.Default()

	// Register API routes.
	routes.SetupRoutes(router, receiptController)

	// Define the server port.
	port := ":8080"
	fmt.Println("Server is running on port", port)

	// Start the server and listen for requests.
	if err := router.Run(port); err != nil {
		logrus.Fatal("Failed to start server:", err)
	}
}
