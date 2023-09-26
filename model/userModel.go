package model

import (
	"database/sql"
	"fmt"

	"github.com/prajwalad101/datekeeper/datastore"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password"`
}

// Password field is ignored
type UserResponse struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"-"`
}

func CreateUser(user *User) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	query := `INSERT INTO users (id, first_name, last_name, email, password) 
    VALUES (DEFAULT, $1, $2, $3, $4)`

	_, err = datastore.DB.Exec(query, user.FirstName, user.LastName, user.Email, passwordHash)
	return err
}

func GetUserByID(id int) (*UserResponse, error) {
	if id == 0 {
		return nil, fmt.Errorf("user id is required")
	}
	query := `SELECT first_name, last_name, email FROM users WHERE id=$1`
	row := datastore.DB.QueryRow(query, id)

	user := new(UserResponse)
	err := row.Scan(&user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(email string) (*UserResponse, error) {
	if email == "" {
		return nil, fmt.Errorf("user email is required")
	}
	query := `SELECT id, first_name, last_name, email, password FROM users WHERE email=$1`
	row := datastore.DB.QueryRow(query, email)

	user := new(UserResponse)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("User not found")
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
