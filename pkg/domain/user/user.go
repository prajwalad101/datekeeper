package user

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Prajwalad101/datekeeper/pkg/datastore"
)

type User struct {
	userName    string
	fullName    string
	email       string
	phoneNumber string
	password    string
	userType    string
}

func ListUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := datastore.DB.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.userName, &user.fullName, &user.password)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, user := range users {
		fmt.Fprintf(w, "%s, %s, %s \n", user.userName, user.fullName, user.password)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	userName := r.FormValue("userName")
	if userName == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	row := datastore.DB.QueryRow("SELECT * FROM users where userName = ?", userName)

	user := new(User)
	err := row.Scan(&user.userName, &user.fullName, &user.password)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s \n", user.userName, user.fullName, user.password)
}
