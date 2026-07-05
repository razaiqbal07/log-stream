package inmemory

import "github.com/razaiqbal07/log-stream/server/internal/model"

type LogRepository struct {
	logs []model.Log
}

func NewLogRepository() *LogRepository {
	return &LogRepository{
		logs: make([]model.Log, 0),
	}
}

func (r *LogRepository) CreateLog(log model.Log) error {
	r.logs = append(r.logs, log)
	return nil
}

func (r *LogRepository) GetLogs() []model.Log {
	return r.logs
}
