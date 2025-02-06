package main

import (
	"fmt"
	"net/http"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Staring Rest API...")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /receipts/process", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Process")
	})

	mux.HandleFunc("GET /receipts/{id}/points", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		uuid := uuid.New()
		fmt.Println(uuid)
		fmt.Fprintf(w, "Points %s", id)
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err)
	}
}