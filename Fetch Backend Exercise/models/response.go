// models/response.go
package models

// ProcessReceiptResponse is the response for processing a receipt
type ProcessReceiptResponse struct {
	ID string `json:"id"`
}

// GetPointsResponse is the response for getting points for a receipt
type GetPointsResponse struct {
	Points int `json:"points"`
}
