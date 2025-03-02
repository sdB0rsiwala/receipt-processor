package utils

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"receipt-processor/models"
)

// points_calculator.go
// This file implements the rules for calculating points based on a receipt's details.
// It includes multiple functions that apply different point rules and return the total score.

// CalculatePoints applies all scoring rules to a receipt and returns the total points.
func CalculatePoints(receipt *models.Receipt) int {
	totalPoints := 0

	// Rule 1: One point for every alphanumeric character in the retailer name.
	totalPoints += countAlphanumeric(receipt.Retailer)

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	if isRoundDollar(receipt.Total) {
		totalPoints += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	if isMultipleOfQuarter(receipt.Total) {
		totalPoints += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	totalPoints += (len(receipt.Items) / 2) * 5

	// Rule 5: Extra points for items where description length is a multiple of 3.
	for _, item := range receipt.Items {
		totalPoints += calculateItemPoints(item)
	}

	// Rule 6: 6 points if the purchase date's day is odd.
	if isOddDay(receipt.PurchaseDate) {
		totalPoints += 6
	}

	// Rule 7: 10 points if the purchase time is between 2:00 PM and 4:00 PM.
	if isAfternoonPurchase(receipt.PurchaseTime) {
		totalPoints += 10
	}

	return totalPoints
}

// countAlphanumeric counts the number of alphanumeric characters in a given string.
func countAlphanumeric(s string) int {
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	return len(re.FindAllString(s, -1))
}

// isRoundDollar checks if the total amount is a whole number without cents.
func isRoundDollar(total string) bool {
	val, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}
	return val == math.Floor(val)
}

// isMultipleOfQuarter checks if the total is a multiple of 0.25.
func isMultipleOfQuarter(total string) bool {
	val, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}
	return math.Mod(val, 0.25) == 0
}

// calculateItemPoints awards points if an item's description length is a multiple of 3.
func calculateItemPoints(item models.Item) int {
	trimmedDesc := strings.TrimSpace(item.ShortDescription)
	if len(trimmedDesc)%3 == 0 {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return 0
		}
		return int(math.Ceil(price * 0.2)) // Multiply price by 0.2 and round up
	}
	return 0
}

// isOddDay returns true if the day of the purchase date is an odd number.
func isOddDay(dateStr string) bool {
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return false
	}
	return parsedDate.Day()%2 != 0
}

// isAfternoonPurchase checks if the purchase time is between 2:00 PM and 4:00 PM.
func isAfternoonPurchase(timeStr string) bool {
	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return false
	}
	hour, minute := parsedTime.Hour(), parsedTime.Minute()
	return (hour == 14) || (hour == 15 && minute == 0)
}
