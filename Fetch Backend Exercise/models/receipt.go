package models

// Receipt is the model for a receipt
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item is the model for an item in a receipt
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
