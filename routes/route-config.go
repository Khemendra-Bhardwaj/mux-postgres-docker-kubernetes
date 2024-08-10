package routes

import (
	// "backend/consumehandlers"
	consumehandlers "backend/pulsarutils/consume-handlers"
	"backend/routes/handler"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	router := mux.NewRouter()
	// TODO: improve api naming and structuring

	router.HandleFunc("/get-users", handler.GetEmployees).Methods("GET")
	router.HandleFunc("/create-user", handler.CreateEmployee).Methods("POST")
	router.HandleFunc("/get-department", handler.GetDepartments).Methods("GET")
	router.HandleFunc("/create-department", handler.CreateDepartment).Methods("POST")

	/*Api for testing Employee logs    */
	router.HandleFunc("/consumeEmployeeLogs", consumehandlers.ConsumeEmployeeLogs).Methods("GET")
	router.HandleFunc("/consumeDepartmentLogs", consumehandlers.ConsumeDepartmentLogs).Methods("GET")
	return router

}
