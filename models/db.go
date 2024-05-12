package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	dbURL := os.Getenv("DB_URL")

	sdb, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	err = sdb.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db = sdb
}

func Close() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
