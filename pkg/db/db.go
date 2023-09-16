package db

import (
	"database/sql"
	"fmt"

	"github.com/prajwalad101/datekeeper/pkg/utils"
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
	env := utils.GetEnv()
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
	)
}
