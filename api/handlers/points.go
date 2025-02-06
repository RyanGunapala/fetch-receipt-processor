package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Points struct {
	Points int `json:"points"`
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Println(id)

	// TODO: Implement the logic to get the points for the given receipt ID
	points := Points{0}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(points)
}
