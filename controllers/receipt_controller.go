package controller

import (
	"net/http"
	"receipt-processor/models"
	"receipt-processor/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// receipt_controller.go
// This file handles HTTP requests for processing receipts and retrieving points.
// It defines the ReceiptController, which interacts with the ReceiptService layer.

// ReceiptController manages API endpoints related to receipts.
type ReceiptController struct {
	ReceiptService *services.ReceiptService // Service layer for business logic
}

// NewReceiptController initializes a new ReceiptController.
func NewReceiptController(receiptService *services.ReceiptService) *ReceiptController {
	return &ReceiptController{ReceiptService: receiptService}
}

// ProcessReceipt handles POST requests to store a new receipt.
// It validates the input, processes the receipt, and returns a unique ID.
func (rc *ReceiptController) ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt

	// Parse and bind the JSON request body to the Receipt struct.
	if err := c.ShouldBindJSON(&receipt); err != nil {
		logrus.WithError(err).Error("Invalid request payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Validate receipt fields before processing.
	if err := receipt.Validate(); err != nil {
		logrus.WithError(err).Error("Invalid receipt data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process the receipt and retrieve its unique ID.
	id := rc.ReceiptService.ProcessReceipt(&receipt)

	logrus.WithFields(logrus.Fields{"id": id}).Info("Receipt processed successfully")

	// Respond with the generated receipt ID.
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetPoints handles GET requests to retrieve points for a given receipt ID.
// If the receipt exists, it returns the calculated points; otherwise, it returns a 404 error.
func (rc *ReceiptController) GetPoints(c *gin.Context) {
	id := c.Param("id") // Extract receipt ID from URL path

	// Fetch points associated with the receipt ID.
	points, found := rc.ReceiptService.GetPoints(id)

	if !found {
		logrus.WithFields(logrus.Fields{"id": id}).Warn("Receipt not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that ID"})
		return
	}

	logrus.WithFields(logrus.Fields{"id": id}).Info("Points retrieved successfully")

	// Respond with the calculated points.
	c.JSON(http.StatusOK, gin.H{"points": points})
}
