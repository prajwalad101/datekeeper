package main

import (
	"log"
	"net/http"

	"github.com/Prajwalad101/datekeeper/middleware"
	"github.com/Prajwalad101/datekeeper/pkg/datastore"
	"github.com/Prajwalad101/datekeeper/pkg/db"
	"github.com/Prajwalad101/datekeeper/pkg/domain/event"
	"github.com/Prajwalad101/datekeeper/pkg/domain/user"
	_ "github.com/go-sql-driver/mysql"
	"github.com/justinas/alice"
)

/* type User struct {
	Name     string
	Password string
} */

/* func login(response http.ResponseWriter, request *http.Request) {
	var user User

	err := utils.DecodeJSONBody(response, request, &user)
	if err != nil {
		var mr *utils.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(response, mr.Msg, mr.Status)
		} else {
			log.Print(err.Error())
			http.Error(response,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
		}
		return
	}

	fmt.Fprintf(response, "Person: %+v", user)
	// if true, create a jwt token and return it to the user
	var (
		key   []byte
		token *jwt.Token
		s     string
	)

	key = []byte(os.Getenv("JWT_SECRET"))

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "my-auth-server",
		"sub": "prajwal",
	})

	s, _ = token.SignedString(key)
	fmt.Println(s)
} */

func init() {
	db := db.InitDB()
	datastore.DB = db
}

func main() {
	mux := http.NewServeMux()

	createEventHandler := http.HandlerFunc(event.CreateEvent)

	mux.HandleFunc("/users", user.ListUser)
	mux.HandleFunc("/user", user.GetUser)

	// create a reusable middleware chain
	stdChain := alice.New(middleware.EnforceJSONHandler)

	mux.Handle("/event", stdChain.Then(createEventHandler))

	log.Print("Listening on :8080")
	httpErr := http.ListenAndServe(":8080", mux)
	log.Fatal(httpErr)
}
