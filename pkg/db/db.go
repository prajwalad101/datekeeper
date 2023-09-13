package db

import (
	"database/sql"
	"fmt"
	"os"
)

// InitDB initialises database session
func InitDB() *sql.DB {
	db, err := sql.Open("postgres", generatePgConnectionString())
	if err != nil {
		panic(fmt.Errorf("[ERROR] %v", err))
	}
	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("[ERROR] %v", err))
	}
	fmt.Println("Connection Successful...")
	return db
}

func generatePgConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
}
