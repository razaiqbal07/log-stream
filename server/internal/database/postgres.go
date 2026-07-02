package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() *sql.DB {
	connStr := os.Getenv("DB_URL")

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		//
	}
	err = db.Ping()
	if err != nil {
		//
	}

	fmt.Println("Connected to db successfully")
	return db

}
