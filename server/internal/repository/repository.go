package repository

import "github.com/razaiqbal07/log-stream/server/internal/model"

type LogRepository interface {
	CreateLog(log model.Log) error
	GetLogs() []model.Log
}
