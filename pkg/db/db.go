package db

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
)

// Initializes database session
func InitDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", "friday@unix(/var/run/mysqld/mysqld.sock)/datekeeper_db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
