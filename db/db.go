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
	DB, ErrorDB = sql.Open("postgres", "postgres://jyazbeck:P@$$w0rd@193.227.182.203:5432/aswaq?sslmode=disable")
	if ErrorDB != nil {
		return ErrorDB, false
	}
	return nil, true
}
