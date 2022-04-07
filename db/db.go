package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	DB      *sql.DB
	ErrorDB error
)

func InitializeDB() (error, bool) {
	DB, ErrorDB = sql.Open("postgres", "postgres://postgres:2400@localhost/stores?sslmode=disable")
	if ErrorDB != nil {
		return ErrorDB, false
	}
	return nil, true
}
