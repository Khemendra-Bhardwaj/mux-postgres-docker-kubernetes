package routes

import (
	// "backend/consumehandlers"
	consumehandlers "backend/pulsarutils/consume-handlers"
	"backend/routes/handler"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	router := mux.NewRouter()
	// TODO: improve api naming

	router.HandleFunc("/get-users", handler.GetEmployees).Methods("GET")
	router.HandleFunc("/create-user", handler.CreateEmployee).Methods("POST")
	router.HandleFunc("/get-department", handler.GetDepartments).Methods("GET")
	router.HandleFunc("/create-department", handler.CreateDepartment).Methods("POST")

	router.HandleFunc("/get-userwithdepartment", handler.GetUsersWithDepartments).Methods("GET")
	router.HandleFunc("/get-departments-users", handler.GetDepartmentEmployees).Methods("GET")

	/*Api for testing Employee logs    */
	router.HandleFunc("/consumeEmployeeLogs", consumehandlers.ConsumeLogs).Methods("GET")

	return router

}
