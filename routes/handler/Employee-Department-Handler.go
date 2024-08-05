package handler

import (
	"backend/db"
	"backend/db/queries"

	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
)

func GetUsersWithDepartments(writer http.ResponseWriter, reader *http.Request) {

	dbConn := db.SetupConnectionDB()
	transaction, err := dbConn.Begin()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer transaction.Rollback()

	rows, err := transaction.Query(queries.GetEmployeesWithDepartments)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var EmployeeDepartmentArr []db.EmployeeDepartment

	for rows.Next() {
		var employeeID int
		var employeeName string
		var departmentID int
		var departmentName string

		err := rows.Scan(&employeeID, &employeeName, &departmentID, &departmentName)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		emp1 := db.EmployeeDepartment{
			EmployeeId:     employeeID,
			EmployeeName:   employeeName,
			DepartmentID:   departmentID,
			DepartmentName: departmentName,
		}

		EmployeeDepartmentArr = append(EmployeeDepartmentArr, emp1)

	}

	if err := rows.Err(); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := transaction.Commit(); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(writer).Encode(&EmployeeDepartmentArr)

}
