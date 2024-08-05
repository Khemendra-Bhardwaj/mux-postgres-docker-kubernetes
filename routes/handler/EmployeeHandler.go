package handler

import (
	"backend/db"
	"backend/db/queries"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func GetEmployees(writer http.ResponseWriter, reader *http.Request) {
	// var dbConn *sql.DB
	dbConn := db.SetupConnectionDB()

	transaction, err := dbConn.Begin()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer transaction.Rollback()

	rows, err := transaction.Query(queries.GetEmployees)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var Employees []map[string]interface{}

	for rows.Next() {
		var employeeID int
		var employeeName string
		var departmentID sql.NullInt32

		if err := rows.Scan(&employeeID, &employeeName, &departmentID); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}

		Employees = append(Employees, map[string]interface{}{
			"employee_id":   employeeID,
			"employee_name": employeeName,
			"department_id": departmentID.Int32,
		})

	}

	if err := transaction.Commit(); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(&Employees); err != nil {
		log.Printf("Error encoding response: %v", err)
	}

}

func CreateEmployee(writer http.ResponseWriter, reader *http.Request) {
	var Employee db.User
	if err := json.NewDecoder(reader.Body).Decode(&Employee); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	dbConn := db.SetupConnectionDB()

	_, err := dbConn.Exec(queries.CreateEmployee, Employee.EmployeeName, Employee.DepartmentID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
