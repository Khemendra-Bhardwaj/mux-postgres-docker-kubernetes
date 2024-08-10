package handler

import (
	"backend/db"
	"backend/db/models"
	"backend/pulsarutils"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/apache/pulsar-client-go/pulsar"
)

func GetEmployees(writer http.ResponseWriter, reader *http.Request) {
	var employees []models.Employee
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
	var empReq models.Employee // fetching model-schema from model.go
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

	// Produce a log message
	logMessage := map[string]string{
		"message":  "Employee created",
		"employee": empReq.EmployeeName,
	}
	msgData, _ := json.Marshal(logMessage)
	_, err := pulsarutils.LogsEmployeeProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})

	if err != nil {
		log.Printf("Error sending create employee log message to Pulsar: %v", err)
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Employee created successfully"))
}
