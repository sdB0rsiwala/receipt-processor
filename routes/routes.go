package routes

import (
	controller "receipt-processor/controllers"

	"github.com/gin-gonic/gin"
)

// routes.go
// This file defines the API routes and maps them to their respective controller functions.
// It ensures that HTTP requests are routed to the correct handlers for processing receipts and retrieving points.

// SetupRoutes registers API endpoints and their corresponding handlers.
func SetupRoutes(router *gin.Engine, receiptController *controller.ReceiptController) {
	// Route to process a receipt and store it in memory.
	// Example: POST /receipts/process
	router.POST("/receipts/process", receiptController.ProcessReceipt)

	// Route to retrieve points for a specific receipt using its ID.
	// Example: GET /receipts/{id}/points
	router.GET("/receipts/:id/points", receiptController.GetPoints)
}
