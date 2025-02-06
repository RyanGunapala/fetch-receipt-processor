package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Id struct {
	Id uuid.UUID `json:"id"`
}

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the logic to process the receipt
	uuid := uuid.New()
	id := Id{uuid}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
