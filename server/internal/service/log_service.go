package service

import (
	"fmt"

	"github.com/razaiqbal07/log-stream/server/internal/model"
	"github.com/razaiqbal07/log-stream/server/internal/repository"
)

type LogService struct {
	logs          []model.Log
	logRepository *repository.LogRepository
}

func NewLogService(logRepository *repository.LogRepository) *LogService {
	return &LogService{
		logRepository: logRepository,
		logs:          make([]model.Log, 0),
	}
}

func (s *LogService) CreateLog(log model.Log) {
	err := s.logRepository.CreateLog(log)
	if err != nil {
		fmt.Println(err)
	}

	s.logs = append(s.logs, log)
}

func (s *LogService) GetLogs() []model.Log {
	logs := s.logRepository.GetLogs()
	return logs
}
