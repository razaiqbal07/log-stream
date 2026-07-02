package service

import "github.com/razaiqbal07/log-stream/server/internal/model"

// type LogService struct {
// 	//
// }

// func NewLogService() *LogService {
// 	return &LogService{}
// }

// func (s *LogService) CreateLog(log model.Log) error {
// 	return nil
// }

var logs []model.Log

func AddLog(log model.Log) {
	logs = append(logs, log)
}

func GetLogs() []model.Log {
	return logs
}
