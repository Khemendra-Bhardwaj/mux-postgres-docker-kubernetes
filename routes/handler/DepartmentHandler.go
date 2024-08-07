package handler

import (
	"backend/db"
	"backend/db/queries"
	"encoding/json"
	"net/http"
)

// GetDepartments retrieves all departments from the database and returns them as JSON.
func GetDepartments(writer http.ResponseWriter, reader *http.Request) {
	var departments []queries.Department
	if err := db.DB.Find(&departments).Error; err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(departments); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

// CreateDepartment creates a new department based on the request payload.
func CreateDepartment(writer http.ResponseWriter, reader *http.Request) {
	var deptReq queries.Department
	if err := json.NewDecoder(reader.Body).Decode(&deptReq); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if deptReq.DepartmentName == "" {
		http.Error(writer, "Department name is required", http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&deptReq).Error; err != nil {
		http.Error(writer, "Error creating department", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Department created successfully"))
}
