package handlers

import (
	"encoding/json"
	"fmt"
	"logstream/server/internal/models"
	"net/http"
)

func HandleIngest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload models.LogPayload
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	fmt.Printf("[INGESTED] [%s] Service: %s | Message: %s\n", payload.Level, payload.Service, payload.Message)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"status":"success"}`))
}
