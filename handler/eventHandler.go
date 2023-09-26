package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/prajwalad101/datekeeper/model"
	"github.com/prajwalad101/datekeeper/utils"
)

func HandleListEvents(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	userID := ctx.Value("userID").(int)

	id := strconv.Itoa(userID)

	events, err := model.GetEvents(id)
	if err != nil {
		return err
	}

	resp := map[string]any{
		"message": "Events fetched successfully",
		"data":    events,
	}

	return utils.WriteJSON(w, http.StatusOK, resp)
}

func HandleGetEvent(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	ctx := r.Context()
	userID := ctx.Value("userID").(int)

	event, err := model.GetEventByID(id, userID)

	if err == sql.ErrNoRows {
		return utils.WriteJSON(
			w,
			http.StatusNotFound,
			utils.ApiError{Error: "Event not found"},
		)
	} else if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, event)
}

func HandleCreateEvent(w http.ResponseWriter, r *http.Request) error {
	e := new(model.Event)

	ctx := r.Context()
	userID := ctx.Value("userID").(int)

	err := utils.DecodeJSONBody(w, r, e)
	if err != nil {
		return err
	}

	if e.Name == "" || e.Date == "" {
		return fmt.Errorf("Please provide required fields (name,  date)")
	}

	err = model.CreateEvent(e, userID)
	if err != nil {
		return err
	}

	resp := utils.JSONResponse{Status: http.StatusOK, Message: "Event created successfully"}
	return utils.WriteJSON(w, http.StatusOK, resp)
}

func HandleUpdateEvent(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	ctx := r.Context()
	userID := ctx.Value("userID").(int)

	e := new(model.Event)
	err := utils.DecodeJSONBody(w, r, e)
	if err != nil {
		return err
	}

	err = model.UpdateEvent(id, userID, e)
	if err != nil {
		return err
	}

	response := map[string]any{
		"status":  http.StatusOK,
		"message": "Successfully updated event",
	}

	return utils.WriteJSON(w, http.StatusOK, response)
}

func HandleDeleteEvent(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	ctx := r.Context()
	userID := ctx.Value("userID").(int)

	err := model.DeleteEvent(id, userID)
	if err != nil {
		return err
	}

	resp := utils.JSONResponse{
		Status:  http.StatusOK,
		Message: "Successfully deleted event",
	}
	return utils.WriteJSON(w, http.StatusOK, resp)
}
