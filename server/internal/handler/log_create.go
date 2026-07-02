package handler

import (
	"encoding/json"
	"net/http"

	"github.com/razaiqbal07/log-stream/server/internal/model"
	"github.com/razaiqbal07/log-stream/server/internal/service"
)

// var logs []model.Log

// type LogHandler struct {
// 	service *LogService
// }

func LogCreate(w http.ResponseWriter, r *http.Request) {
	var log model.Log
	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	service.AddLog(log)
	// logs = append(logs, log)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "success"})
}
