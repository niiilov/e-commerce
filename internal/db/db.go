package db

import (
	"database/sql"

	"github.com/niiilov/e-commerce/internal/utils"
)

var DB *sql.DB

func InitDB() error {
	psqlInfo, err := utils.Env("DB_CONNECT")
	if err != nil {
		return err
	}

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
