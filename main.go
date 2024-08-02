package main

import (
	"backend/db"
	"backend/routes"

	"net/http"
)

func main() {

	db.SetupDatabase() // TODO: setup locks for queries

	router := routes.SetupRouter()

	http.ListenAndServe(":8080", router)

}
