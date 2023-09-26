package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prajwalad101/datekeeper/model"
	"github.com/prajwalad101/datekeeper/utils"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) error {
	user := new(model.User)

	err := utils.DecodeJSONBody(w, r, &user)
	if err != nil {
		return err
	}

	if user.Email == "" || user.Password == "" {
		return fmt.Errorf("Please provide email and password")
	}

	err = model.CreateUser(user)
	if err != nil {
		return err
	}

	resp := utils.JSONResponse{
		Message: "Successfully registered user",
		Status:  http.StatusOK,
	}

	return utils.WriteJSON(w, http.StatusOK, resp)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	input := new(model.User)

	err := utils.DecodeJSONBody(w, r, &input)
	if err != nil {
		return err
	}

	if input.Password == "" {
		return fmt.Errorf("Password is required")
	}

	// check if the user exists
	user, err := model.GetUserByEmail(input.Email)
	if err != nil {
		return err
	}

	// check if the password hash matched the provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return fmt.Errorf("Invalid credentials")
	}

	token, err := utils.CreateAccessToken(user.ID, time.Now())
	if err != nil {
		return err
	}

	resp := map[string]any{
		"message": "Sucessfully logged in",
		"data": map[string]any{
			"user":        user,
			"accessToken": token,
		},
	}

	return utils.WriteJSON(w, http.StatusOK, resp)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	id := ctx.Value("userID").(int)

	user, err := model.GetUserByID(id)
	if err != nil {
		return err
	}

	resp := map[string]any{
		"message": "User details fetched successfully",
		"data":    user,
	}

	return utils.WriteJSON(w, http.StatusOK, resp)
}
