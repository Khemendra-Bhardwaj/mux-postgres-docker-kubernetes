package handler

import (
	"backend/db"
	"backend/db/queries"
	"encoding/json"
	"net/http"
)

func GetEmployees(writer http.ResponseWriter, reader *http.Request) {
	var employees []queries.Employee
	if err := db.DB.Preload("Department").Find(&employees).Error; err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(employees); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func CreateEmployee(writer http.ResponseWriter, reader *http.Request) {
	var empReq queries.Employee
	if err := json.NewDecoder(reader.Body).Decode(&empReq); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if empReq.EmployeeName == "" {
		http.Error(writer, "Employee name is required", http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&empReq).Error; err != nil {
		http.Error(writer, "Error creating employee", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Employee created successfully"))
}
