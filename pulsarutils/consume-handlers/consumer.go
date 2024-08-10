package consumehandlers

import (
	"backend/pulsarutils"
	"context"
	"log"
	"net/http"
)

// TODO : implement for separate Department Logs as Well

func ConsumeEmployeeLogs(w http.ResponseWriter, r *http.Request) {
	// for {
	msg, err := pulsarutils.LogsEmployeeConsumer.Receive(context.Background())
	if err != nil {
		log.Printf("Error receivinglogs Employees : %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	log.Printf("Received log Employee message: %s", string(msg.Payload()))
	pulsarutils.LogsEmployeeConsumer.Ack(msg)
}

func ConsumeDepartmentLogs(w http.ResponseWriter, r *http.Request) {
	msg, err := pulsarutils.LogsDepartmentConsumer.Receive(context.Background())

	if err != nil {
		log.Printf("Error receiving logs Department : %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("Received log Department message: %s", string(msg.Payload()))
	pulsarutils.LogsDepartmentConsumer.Ack(msg)

}
