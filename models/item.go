package models

// item.go
// This file defines the Item model, representing an individual product
// on a receipt, including its description and price.

// Item represents a product in the receipt with a description and price.
type Item struct {
	ShortDescription string `json:"shortDescription"` // Name or description of the item
	Price            string `json:"price"`            // Price of the item as a string (e.g., "6.49")
}
