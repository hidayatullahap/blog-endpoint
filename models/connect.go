package models

import (
	"os"

	"github.com/jmoiron/sqlx"
)

func DbConnect() (db *sqlx.DB) {
	dbName := os.Getenv("DATABASE_NAME")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbConnect := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	dbConnection := sqlx.MustConnect("mysql", dbConnect)

	return dbConnection
}
