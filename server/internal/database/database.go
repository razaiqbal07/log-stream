package database

import (
	"github.com/razaiqbal07/log-stream/server/internal/config"
)

func DatabaseResolver(config *config.AppConfig) any {
	switch config.DATABASE {
	case "postgres":
		return Connect()
	default:
		return nil
	}
}
