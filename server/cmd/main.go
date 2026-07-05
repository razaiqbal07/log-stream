package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/razaiqbal07/log-stream/server/internal/handler"
	"github.com/razaiqbal07/log-stream/server/internal/repository/inmemory"
	"github.com/razaiqbal07/log-stream/server/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		// log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	// db := database.Connect()

	// logRepository := postgres.NewLogRepository(db)
	lr := inmemory.NewLogRepository()
	logService := service.NewLogService(lr)
	logHandler := handler.NewLogHandler(logService)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", health)

	/** Routes*/
	mux.HandleFunc("POST /api/logs", logHandler.LogCreate)
	mux.HandleFunc("GET /api/logs", logHandler.LogRead)

	http.ListenAndServe(PORT, mux)
}

func health(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Server running..."}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
