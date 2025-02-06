package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/ryangunapala/fetch-receipt-processor/internal/models"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log the receipt
	log.Println(receipt)

	// TODO: Implement the logic to process the receipt
	uuid := uuid.New()
	id := models.Id{Id: uuid}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
