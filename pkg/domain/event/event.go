package event

import (
	"fmt"
	"net/http"

	"github.com/Prajwalad101/datekeeper/pkg/utils"
)

type Event struct {
	Name string
	Date string
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event

	// get the token from authorization header

	utils.DecodeJSONBody(w, r, &event)

	fmt.Println(event)
}
