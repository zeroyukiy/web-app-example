package database

import (
	"database/sql"
	"log"
	"time"
)

var db *sql.DB

func StartConnection() {
	database, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/web-app-example")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// See "Important settings" section.
	database.SetConnMaxLifetime(time.Minute * 1)
	database.SetMaxOpenConns(100)
	database.SetMaxIdleConns(10)

	db = database
}

func GetConnection() *sql.DB {
	return db
}
