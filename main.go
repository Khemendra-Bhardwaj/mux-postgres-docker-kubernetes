package main

import (
	"backend/db"
	"backend/pulsarutils"
	"backend/routes"
	"log"

	"net/http"
)

func main() {

	db.SetupDatabase() // TODO: setup locks for queries

	pulsarutils.SetupPulsar()
	defer pulsarutils.Close()

	log.Println("Connected to DB and Pulsar Done!")

	router := routes.SetupRouter()

	http.ListenAndServe(":8080", router)

}

// build docker image , push it dockerhub and applt kubernetes to it
