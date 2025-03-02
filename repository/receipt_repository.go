package repository

import (
	"receipt-processor/models"
	"sync"

	"github.com/sirupsen/logrus"
)

// receipt_repository.go
// This file manages the in-memory storage of receipts using a map.
// It provides functions to store and retrieve receipts in a thread-safe manner.

// ReceiptRepository provides an in-memory storage for receipts.
type ReceiptRepository struct {
	mu       sync.Mutex                 // Mutex to ensure thread-safe access
	Receipts map[string]*models.Receipt // Map storing receipts with unique IDs as keys
}

// NewReceiptRepository initializes a new receipt repository.
func NewReceiptRepository() *ReceiptRepository {
	return &ReceiptRepository{
		Receipts: make(map[string]*models.Receipt), // Initializes the receipt storage
	}
}

// SaveReceipt stores a receipt in the repository using a unique ID.
func (r *ReceiptRepository) SaveReceipt(id string, receipt *models.Receipt) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Receipts[id] = receipt // Store the receipt in memory

	logrus.WithFields(logrus.Fields{
		"id":             id,
		"stored_receipt": receipt,
	}).Info("Receipt stored successfully") // Log successful storage
}

// GetReceipt retrieves a receipt from the repository by ID.
// Returns the receipt and a boolean indicating whether it exists.
func (r *ReceiptRepository) GetReceipt(id string) (*models.Receipt, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	receipt, exists := r.Receipts[id] // Retrieve receipt from storage
	return receipt, exists
}
