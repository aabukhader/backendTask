package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

// ConnectDb : create the main connection to the database
func ConnectDb() *sql.DB {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/backend")

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
