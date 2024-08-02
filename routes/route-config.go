package routes

import (
	"backend/routes/handler"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/get-users", handler.GetEmployees).Methods("GET")
	router.HandleFunc("/create-user", handler.CreateEmployee).Methods("POST")
	router.HandleFunc("/get-department", handler.GetDepartments).Methods("GET")
	router.HandleFunc("/create-department", handler.CreateDepartment).Methods("POST")

	return router

}
