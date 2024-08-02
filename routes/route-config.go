package routes

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/get-users", GetEmployees).Methods("GET")
	router.HandleFunc("/create-user", CreateEmployee).Methods("POST")
	router.HandleFunc("/get-department", GetDepartments).Methods("GET")
	router.HandleFunc("/create-department", CreateDepartment).Methods("POST")

	return router

}
