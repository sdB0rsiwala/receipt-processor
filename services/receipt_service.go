package services

import (
	"receipt-processor/models"
	"receipt-processor/repository"
	"receipt-processor/utils"
	"sync"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// receipt_service.go
// This file implements the business logic for processing receipts and calculating points.
// It interacts with the repository layer to store and retrieve receipts and applies business rules for points calculation.

// ReceiptService provides methods for processing receipts and retrieving points.
type ReceiptService struct {
	Repo *repository.ReceiptRepository // Repository for storing and retrieving receipts
	mu   sync.Mutex                    // Mutex for thread-safe operations
}

// NewReceiptService initializes a new ReceiptService.
func NewReceiptService(repo *repository.ReceiptRepository) *ReceiptService {
	return &ReceiptService{Repo: repo}
}

// ProcessReceipt generates a unique ID for the receipt, stores it, and returns the ID.
func (s *ReceiptService) ProcessReceipt(receipt *models.Receipt) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New().String() // Generate a unique identifier for the receipt

	s.Repo.SaveReceipt(id, receipt) // Store the receipt in memory

	logrus.WithFields(logrus.Fields{
		"id":       id,
		"retailer": receipt.Retailer,
	}).Info("Receipt processed successfully")

	return id // Return the unique receipt ID
}

// GetPoints retrieves the points awarded for a given receipt ID.
// Returns the total points and a boolean indicating if the receipt exists.
func (s *ReceiptService) GetPoints(id string) (int, bool) {
	receipt, exists := s.Repo.GetReceipt(id) // Retrieve receipt from storage
	if !exists {
		logrus.WithFields(logrus.Fields{"id": id}).Warn("Receipt not found")
		return 0, false // Return 0 points if the receipt does not exist
	}

	points := utils.CalculatePoints(receipt) // Calculate points based on rules

	logrus.WithFields(logrus.Fields{
		"id":     id,
		"points": points,
	}).Info("Points calculated successfully")

	return points, true
}
