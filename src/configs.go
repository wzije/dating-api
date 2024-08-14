package src

import (
	"database/sql"
	"os"
)

func DB() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dbURL)

	if err != nil {
		panic(err.Error())
	}

	return db
}
