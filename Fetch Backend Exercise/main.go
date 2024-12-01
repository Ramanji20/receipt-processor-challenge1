// Package main provides the main implementation of a receipt processing service
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"receipt-processor-challenge/models"
	"receipt-processor-challenge/utils"
)

// Use mutex to protect shared data
var (
	mu        sync.RWMutex
	receipts  = make(map[string]models.Receipt)
	pointsMap = make(map[string]int)
)

// Server configuration constants
const (
	serverPort = ":8081"
)

func main() {
	// Initialize router
	router := initRouter()

	// Start server
	log.Printf("Server started on port %s\n", serverPort)
	if err := http.ListenAndServe(serverPort, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// initRouter initializes and configures routes
func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", processReceipt).Methods(http.MethodPost)
	r.HandleFunc("/receipts/{id}/points", getPoints).Methods(http.MethodGet)
	return r
}

// processReceipt processes receipts and calculates points
// @Summary Process receipt information
// @Description Accepts receipt info, validates format, calculates points and returns ID
// @Accept json
// @Produce json
// @Success 200 {object} models.ProcessReceiptResponse
// @Failure 400 {string} string "Invalid receipt format"
func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt

	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid receipt format", http.StatusBadRequest)
		return
	}

	// Validate receipt format
	if ok := utils.ValidateReceipt(receipt); !ok {
		http.Error(w, "Invalid receipt format", http.StatusBadRequest)
		return
	}

	// Generate ID and save data
	id := uuid.New().String()
	points := utils.CalculatePoints(receipt)

	mu.Lock()
	receipts[id] = receipt
	pointsMap[id] = points
	mu.Unlock()

	// Return response
	w.Header().Set("Content-Type", "application/json")
	response := models.ProcessReceiptResponse{ID: id}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Response encoding error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// getPoints gets points for a receipt by ID
// @Summary Get receipt points
// @Description Returns points for a receipt based on ID
// @Produce json
// @Param id path string true "Receipt ID"
// @Success 200 {object} models.GetPointsResponse
// @Failure 404 {string} string "Receipt not found"
func getPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Query points
	mu.RLock()
	points, exists := pointsMap[id]
	mu.RUnlock()

	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	response := models.GetPointsResponse{Points: points}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Response encoding error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
