package handler

import (
	"encoding/json"
	"net/http"

	"github.com/razaiqbal07/log-stream/server/internal/model"
)

func (h *Handler) LogCreate(w http.ResponseWriter, r *http.Request) {
	var log model.Log

	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	h.logService.CreateLog(log)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "success"})
}
