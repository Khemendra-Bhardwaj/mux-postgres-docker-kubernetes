package pulsarutils

import (
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
	// Replace with your actual module path
)

var Client pulsar.Client

// var Producer pulsar.Producer
// var Consumer pulsar.Consumer
var LogsProducer pulsar.Producer
var LogsConsumer pulsar.Consumer

func SetupPulsar() {
	var err error
	Client, err = pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://10.107.11.139:6650", // Replace with your Pulsar broker URL
	})
	if err != nil {
		log.Fatalf("Could not create Pulsar client: %v", err)
	}
	log.Println("Created  a pulsar client ")
	// Producer, err = Client.CreateProducer(pulsar.ProducerOptions{
	// 	Topic: "employee-topic", // Replace with your topic name
	// })

	LogsProducer, err = Client.CreateProducer(pulsar.ProducerOptions{
		Topic: "employee-logs",
	})
	if err != nil {
		log.Println("Error Creating Employee Producer Logs ")
	}
	log.Println("Created Employee Producer logs success")

	LogsConsumer, err = Client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "employee-logs",
		SubscriptionName: "logs-subscription",
	})
	if err != nil {
		log.Println("Error Creating Employee Logs Consumer ")
	}
	log.Println("Created Employee Logs Consumer success")

	// if err != nil {
	// 	log.Fatalf("Could not create Pulsar producer: %v", err)
	// }
	// log.Println("Created  a producer  ")

	// Consumer, err = Client.Subscribe(pulsar.ConsumerOptions{
	// 	Topic:            "employee-topic",        // Replace with your topic name
	// 	SubscriptionName: "employee-subscription", // Replace with your subscription name
	// })
	// if err != nil {
	// 	log.Fatalf("Could not create Pulsar consumer: %v", err)
	// }

	// log.Println("Created  a Consumer   ")
}

func Close() {
	if LogsProducer != nil {
		LogsProducer.Close()
	}

	if LogsConsumer != nil {
		LogsConsumer.Close()
	}
	if Client != nil {
		Client.Close()
	}
}
