package utils

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"receipt-processor-challenge/models"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// One point for every alphanumeric character in the retailer name
	alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(alphanumeric.FindAllString(receipt.Retailer, -1))

	// 50 points if the total is a round dollar amount with no cents
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// Points based on item descriptions
	for _, item := range receipt.Items {
		description := strings.TrimSpace(item.ShortDescription)
		if len(description) % 3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day() % 2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)

	// purchase time should be from 14:01 to 15:59
	if (purchaseTime.Hour() == 14 && purchaseTime.Minute() > 0) || (purchaseTime.Hour() == 15 && purchaseTime.Minute() < 60) {
		points += 10
	}

	return points
}
