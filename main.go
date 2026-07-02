package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	mux.HandleFunc("POST /api/log", injestHandler)
	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "OK"})
}

func injestHandler(w http.ResponseWriter, r *http.Request) {

}
