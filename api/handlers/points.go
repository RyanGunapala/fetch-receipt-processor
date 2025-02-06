package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ryangunapala/fetch-receipt-processor/internal/models"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Println(id)

	// TODO: Implement the logic to get the points for the given receipt ID
	points := models.Points{Points: 0}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(points)
}
