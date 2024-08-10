package pulsarutils

import (
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
	// Replace with your actual module path
)

var Client pulsar.Client

var LogsEmployeeProducer pulsar.Producer
var LogsEmployeeConsumer pulsar.Consumer
var LogsDepartmentProducer pulsar.Producer
var LogsDepartmentConsumer pulsar.Consumer

func SetupPulsar() {
	var err error
	Client, err = pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://10.107.11.139:6650", // cluster ip  of svc pulsar mini proxy
	})
	if err != nil {
		log.Fatalf("Could not create Pulsar client: %v", err)
	}
	log.Println("Created  a pulsar client ")

	/*Employee Logs Start */
	LogsEmployeeProducer, err = Client.CreateProducer(pulsar.ProducerOptions{
		Topic: "employee-logs", // TODO : import from somewhere than hardcoding
	})
	if err != nil {
		log.Println("Error Creating Employee Producer Logs ")
	}
	log.Println("Created Employee Producer logs success")

	LogsEmployeeConsumer, err = Client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "employee-logs",
		SubscriptionName: "logs-employees",
	})

	if err != nil {
		log.Println("Error Creating Employee Logs Consumer ")
	}
	log.Println("Created Employee Logs Consumer success")

	/*Employee Logs End */

	/*Department Logs Start*/

	LogsDepartmentProducer, err = Client.CreateProducer(pulsar.ProducerOptions{
		Topic: "department-logs",
	})
	if err != nil {
		log.Println("Error Creating Department logs")
	}

	log.Println("Created Department Producer logs success")

	LogsDepartmentConsumer, err = Client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "department-logs",
		SubscriptionName: "logs-department",
	})
	if err != nil {
		log.Println("Erorr Creating Department Consumer ")
	}

	/*Department Logs end */

}

func Close() {
	if LogsEmployeeProducer != nil {
		LogsEmployeeProducer.Close()
	}

	if LogsEmployeeConsumer != nil {
		LogsEmployeeConsumer.Close()
	}

	if LogsDepartmentProducer != nil {
		LogsDepartmentProducer.Close()
	}

	if LogsDepartmentConsumer != nil {
		LogsDepartmentConsumer.Close()
	}

	if Client != nil {
		Client.Close()
	}
}
