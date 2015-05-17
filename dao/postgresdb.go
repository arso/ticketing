package dao

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB = connectDB()

func connectDB() *sql.DB {

	db, err := sql.Open("postgres", "user=postgres dbname=ticketing")
	if err != nil {
		log.Fatal(err)
	}

	return db

}
