package db

import (
	"database/sql"
	"todo-app-go/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := config.GetDSN()
	DB, err = sql.Open("mysql", dsn)
	config.CheckErr(err)
}
