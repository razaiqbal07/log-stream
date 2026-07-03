package handler

import "github.com/razaiqbal07/log-stream/server/internal/service"

var logService = &service.LogService{}

type Handler struct {
	logService *service.LogService
}

func NewLogHandler(s *service.LogService) *Handler {
	return &Handler{
		logService: s,
	}
}
