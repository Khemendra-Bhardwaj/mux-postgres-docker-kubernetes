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

// GetDepartments retrieves all departments from the database and returns them as JSON.

// TODO : Logs of Get Department As Well
func GetDepartments(writer http.ResponseWriter, reader *http.Request) {
	var departments []models.Department
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
	var deptReq models.Department
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

	logMessage := map[string]string{
		"message":                 "Department created",
		"department-name-created": deptReq.DepartmentName,
	}

	msgData, _ := json.Marshal(logMessage)

	_, err := pulsarutils.LogsDepartmentProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})

	if err != nil {
		log.Printf("Error sending create department log message to Pulsar %v ", err)
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Department created successfully"))
}
