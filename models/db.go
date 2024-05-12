package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	sdb, err := sql.Open("mysql", "user:pass@/sentinel")
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
