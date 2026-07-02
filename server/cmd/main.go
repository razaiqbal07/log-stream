package main

import (
	"fmt"
	"logstream/server/internal/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/ingest", handlers.HandleIngest)

	fmt.Println("🚀 LogStream Ingestion Engine starting on port 8080...")
	http.ListenAndServe(":8080", mux)
}
