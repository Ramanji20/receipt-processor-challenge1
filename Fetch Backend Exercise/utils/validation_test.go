package utils

import (
	"testing"
	"receipt-processor-challenge/models"
)

func TestReceiptFormat(t *testing.T) {
	tests := []struct {
		name     string
		receipt  models.Receipt
		expected bool
	}{
		{
			name: "Everything is formatted correctly",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.49",
			},
			expected: true,
		},
		{
			name: "Retailer name is not formatted correctly",
			receipt: models.Receipt{
				Retailer:     "()()()()",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.49",
			},
			expected: false,
		},
		{
			name: "Date is not formatted correctly",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-0101",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.49",
			},
			expected: false,
		},
		{
			name: "Hour is not formatted correctly",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "1301",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.49",
			},
			expected: false,
		},
		{
			name: "Total price is not formatted correctly",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.49--+",
			},
			expected: false,
		},
		{
			name: "Item description is not formatted correctly",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "[][][]", Price: "6.49"},
				},
				Total: "6.49",
			},
			expected: false,
		},
		{
			// item price is not formatted correctly
			name: "Item description length multiple of 3",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49+++"},
				},
				Total: "6.49",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateReceipt(tt.receipt); got != tt.expected {
				t.Errorf("validateReceipt() = %v, want %v", got, tt.expected)
			}
		})
	}
}
