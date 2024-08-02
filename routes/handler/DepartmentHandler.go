package handler

import (
	"backend/db"
	"backend/db/queries"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func GetDepartments(writer http.ResponseWriter, reader *http.Request) {

	dbconn := db.SetupConnectionDB()
	transaction, err := dbconn.Begin()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer transaction.Rollback()
	rows, err := transaction.Query(queries.GetDepartments)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var departments []map[string]interface{}

	for rows.Next() {
		var departmentID int
		var departmentName string
		if err := rows.Scan(&departmentID, &departmentName); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}
		departments = append(departments, map[string]interface{}{
			"department_id":   departmentID,
			"department_name": departmentName,
		})
	}

	if err := transaction.Commit(); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(departments); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func CreateDepartment(writer http.ResponseWriter, reader *http.Request) {

	var deptReq db.Department
	if err := json.NewDecoder(reader.Body).Decode(&deptReq); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if deptReq.DepartmentName == "" {
		http.Error(writer, "Department name is required", http.StatusBadRequest)
		return
	}

	dbconn := db.SetupConnectionDB()
	transaction, err := dbconn.Begin()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer transaction.Rollback()

	if dbconn == nil {
		http.Error(writer, "Database connection is not established", http.StatusInternalServerError)
		return
	}

	_, err = transaction.Exec(queries.CreateDepartment, deptReq.DepartmentName)
	if err != nil {
		http.Error(writer, "Error creating department", http.StatusInternalServerError)
		return
	}

	if err := transaction.Commit(); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Department created successfully"))
}
