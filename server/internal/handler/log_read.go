package handler

import (
	"encoding/json"
	"net/http"
)

func LogRead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	logs := logService.GetLogs()

	json.NewEncoder(w).Encode(logs)
}
