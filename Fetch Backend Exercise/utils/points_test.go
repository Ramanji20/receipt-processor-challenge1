package utils

import (
	"testing"
	"receipt-processor-challenge/models"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  models.Receipt
		expected int
	}{
		{
			// Case 1
			name: "Basic case",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expected: 28,
		},
		{
			// Case 2
			name: "Round dollar total and multiple of 0.25",
			receipt: models.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []models.Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				Total: "9.00",
			},
			expected: 109,
		},
		{
			// 5 points - retailer name
			// 10 points - purchased between 2pm and 4pm
			// total 15 points
			name: "Purchase time between 14:01 and 15:59",
			receipt: models.Receipt{
				Retailer:     "Store",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "14:30",
				Items: []models.Item{
					{ShortDescription: "Item1", Price: "1.99"},
				},
				Total: "1.99",
			},
			expected: 15,
		},
		{
			// 5 points - retailer name
			// 0 points - purchased outside 2pm and 4pm
			// total 5 points
			name: "Purchase time outside 14:01 and 15:59",
			receipt: models.Receipt{
				Retailer:     "Store",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "16:00",
				Items: []models.Item{
					{ShortDescription: "Item1", Price: "1.99"},
				},
				Total: "1.99",
			},
			expected: 5,
		},
		{
			// 4 points - retailer name
			// RoundUp(1.2 * 0.2) = 1 point - item description is a multiple of 3
			name: "Item description length multiple of 3",
			receipt: models.Receipt{
				Retailer:     "Shop",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "10:00",
				Items: []models.Item{
					{ShortDescription: "123", Price: "1.2"},
				},
				Total: "1.2",
			},
			expected: 5,
		},
		{
			// case 3
			// 9 points - retailer name
			// 5 points - 2 items
			// RoundUp(1.40 * 0.2) = 1 point - item description is a multiple of 3 (Dasani)
			// total 15
			name: "Morning Receipt",
			receipt: models.Receipt{
				Retailer:     "Walgreens",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "08:13",
				Items: []models.Item{
					{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
					{ShortDescription: "Dasani", Price: "1.40"},
				},
				Total: "2.65",
			},
			expected: 15,
		},
		{
			// case 4
			// 6 points - retailer name
			// 25 points - total is a multiple of 1.25
			// total 31 points
			name: "Simple Receipt",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "13:13",
				Items: []models.Item{
					{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
				},
				Total: "1.25",
			},
			expected: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculatePoints(tt.receipt); got != tt.expected {
				t.Errorf("calculatePoints() = %v, want %v", got, tt.expected)
			}
		})
	}
}
