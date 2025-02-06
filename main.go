package main

import (
	"log"
	"net/http"

	"github.com/ryangunapala/fetch-receipt-processor/api/handlers"
)

func initializeServer() {
	log.Println("Staring the API...")

	http.Handle("GET /receipts/{id}/points", http.HandlerFunc(handlers.GetPoints))
	http.Handle("POST /receipts/process", http.HandlerFunc(handlers.ProcessReceipt))

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func main() {
	initializeServer()
}
