package repository

import (
	"database/sql"

	"github.com/razaiqbal07/log-stream/server/internal/config"
	"github.com/razaiqbal07/log-stream/server/internal/model"
	"github.com/razaiqbal07/log-stream/server/internal/repository/inmemory"
	"github.com/razaiqbal07/log-stream/server/internal/repository/postgres"
)

type LogRepository interface {
	CreateLog(log model.Log) error
	GetLogs() []model.Log
}

func ResolveRepository(appConfig *config.AppConfig, dbType any) LogRepository {
	switch appConfig.DATABASE {
	case "postgres":
		db, ok := dbType.(*sql.DB)
		if !ok {
			panic("Invalid database config")
		}
		return postgres.NewLogRepository(db)
	default:
		return inmemory.NewLogRepository()
	}
}
