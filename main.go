package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ynoacamino/ynoa-shorter/db"
	middleware "github.com/ynoacamino/ynoa-shorter/middlewares"
	"github.com/ynoacamino/ynoa-shorter/routes"
)

func main() {
	db.InitDBConnection()
	defer db.CloseDBConnection()

	app := mux.NewRouter()

	app.Use(middleware.ConncetionSecret)

	shorterRouter := app.PathPrefix("/shorter").Subrouter()

	routes.SetUpShorterRoutes(shorterRouter)

	http.ListenAndServe(":8000", app)

}
