package models

import (
	"github.com/go-playground/validator/v10"
)

// receipt.go
// This file defines the Receipt model, representing a customer's purchase receipt.
// It includes details such as retailer name, purchase date/time, item list, and total amount.
// The file also includes a validation function to ensure correct receipt formatting.

var validate = validator.New() // Validator instance for struct validation

// Receipt represents a customer's purchase receipt.
type Receipt struct {
	Retailer     string `json:"retailer" validate:"required"`     // Name of the retailer/store
	PurchaseDate string `json:"purchaseDate" validate:"required"` // Date of purchase (format: YYYY-MM-DD)
	PurchaseTime string `json:"purchaseTime" validate:"required"` // Time of purchase (format: HH:MM)
	Items        []Item `json:"items" validate:"required,dive"`   // List of purchased items
	Total        string `json:"total" validate:"required"`        // Total amount paid as a string (e.g., "35.35")
}

// Validate checks if the receipt data conforms to the required format.
func (r *Receipt) Validate() error {
	return validate.Struct(r) // Uses go-playground/validator for validation
}
