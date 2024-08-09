package config

import (
	"log"
)

const (
	// Change these values based on your setup
	DBUser     = "root"
	DBPassword = ""
	DBHost     = "localhost"
	DBPort     = "3306"
	DBName     = "todo_db"
)

func GetDSN() string {
	return DBUser + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
