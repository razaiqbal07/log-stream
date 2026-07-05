package postgres

import (
	"database/sql"

	"github.com/razaiqbal07/log-stream/server/internal/model"
)

type LogRepository struct {
	db *sql.DB
}

func NewLogRepository(db *sql.DB) *LogRepository {
	return &LogRepository{
		db: db,
	}
}

func (r *LogRepository) CreateLog(log model.Log) error {
	query := "INSERT INTO logs (message, service, type) values ($1, $2, $3)"
	_, err := r.db.Exec(query, log.Message, log.Service, log.Type)
	return err
}

func (r *LogRepository) GetLogs() []model.Log {
	rows, err := r.db.Query("SELECT service, type, message FROM logs")
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
