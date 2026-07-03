package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) LogRead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	logs := h.logService.GetLogs()

	json.NewEncoder(w).Encode(logs)
}
