package main

import (
	"encoding/json"
	"net/http"

	"github.com/razaiqbal07/log-stream/server/internal/config"
	"github.com/razaiqbal07/log-stream/server/internal/database"
	"github.com/razaiqbal07/log-stream/server/internal/handler"
	"github.com/razaiqbal07/log-stream/server/internal/repository"
	"github.com/razaiqbal07/log-stream/server/internal/service"
)

func main() {

	appConfig := config.ResolveConfig()
	db := database.DatabaseResolver(appConfig)
	repository := repository.ResolveRepository(appConfig, db)

	logService := service.NewLogService(repository)
	logHandler := handler.NewLogHandler(logService)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", health)

	/** Routes*/
	mux.HandleFunc("POST /api/logs", logHandler.LogCreate)
	mux.HandleFunc("GET /api/logs", logHandler.LogRead)

	http.ListenAndServe(appConfig.PORT, mux)
}

func health(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Server running..."}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
