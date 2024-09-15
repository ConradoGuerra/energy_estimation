package main

import (
	"energy_estimation/infrastructure/http_handlers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/api/estimation", http_handlers.EstimationHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)

}
