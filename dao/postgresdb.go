package dao

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB = connectDB()

func connectDB() *sql.DB {

	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB connection established")

	return db

}
