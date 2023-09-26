package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prajwalad101/datekeeper/datastore"
	"github.com/prajwalad101/datekeeper/middleware"
	"github.com/prajwalad101/datekeeper/routes"
	"github.com/prajwalad101/datekeeper/service"
	"github.com/prajwalad101/datekeeper/utils"
)

func init() {
	utils.InitEnv()
	err := datastore.InitDB()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.EnforceJSON)

	apiRouter := chi.NewRouter()
	routes.EventRouter(apiRouter)
	routes.UserRouter(apiRouter)

	r.Mount("/api/v1", apiRouter)

	service.RunSchedule()

	log.Print("Listening on ", utils.Env.Port)
	err := http.ListenAndServe(utils.Env.Port, r)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
}
