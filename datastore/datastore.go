package datastore

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/prajwalad101/datekeeper/utils"
)

var DB *sql.DB

func InitDB() error {
	db, err := sql.Open("postgres", generatePgConnectionString())
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	DB = db
	log.Println("Connection Successful...")
	err = createEventTable()
	err = createUserTable()
	if err != nil {
		return err
	}
	return nil
}

func generatePgConnectionString() string {
	env := utils.Env
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
	)
}

func createEventTable() error {
	query := `CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    user_id integer REFERENCES users,
    name VARCHAR(30) NOT NULL,
    note text,
    date DATE NOT NULL
  )`

	_, err := DB.Exec(query)
	return err
}

func createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    email VARCHAR(30) NOT NULL UNIQUE,
    password VARCHAR(200) NOT NULL
  )`

	_, err := DB.Exec(query)
	return err
}
