package db

import (
	"database/sql"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() error {
	env, err := godotenv.Read(".env")
	if err != nil {
		return err
	}

	psqlInfo := env["DB_CONNECT"]

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}
