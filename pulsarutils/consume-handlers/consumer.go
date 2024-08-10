package consumehandlers

import (
	"backend/pulsarutils"
	"context"
	"log"
	"net/http"
)

// TODO : implement for separate Department Logs as Well
func ConsumeLogs(w http.ResponseWriter, r *http.Request) {
	// for {
	msg, err := pulsarutils.LogsConsumer.Receive(context.Background())
	if err != nil {
		log.Printf("Error receiving message: %v", err)
	}

	log.Printf("Received log message: %s", string(msg.Payload()))
	pulsarutils.LogsConsumer.Ack(msg)
}
