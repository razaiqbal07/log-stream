package service

import (
	"database/sql"
	"fmt"

	"github.com/razaiqbal07/log-stream/server/internal/model"
)

type LogService struct {
	db   *sql.DB
	logs []model.Log
}

func NewLogService(db *sql.DB) *LogService {
	return &LogService{
		db:   db,
		logs: make([]model.Log, 0),
	}
}

func (s *LogService) CreateLog(log model.Log) {
	_, err := s.db.Exec("insert into logs (message, service, type) values ($1, $2, $3)", log.Message, log.Service, log.Type)
	if err != nil {
		fmt.Println(err)
	}

	s.logs = append(s.logs, log)
}

func (s *LogService) GetLogs() []model.Log {
	rows, err := s.db.Query("SELECT service, type, message FROM logs")
	if err != nil {
		// return nil
	}
	defer rows.Close()

	var logs []model.Log

	for rows.Next() {
		var log model.Log

		err := rows.Scan(&log.Service, &log.Type, &log.Message)
		if err != nil {
			// return nil, err
		}

		logs = append(logs, log)
	}

	return logs
}
