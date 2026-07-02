package handler

import (
	"encoding/json"
	"net/http"

	"github.com/razaiqbal07/log-stream/server/internal/service"
)

func LogRead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	logs := service.GetLogs()
	json.NewEncoder(w).Encode(logs)
}
