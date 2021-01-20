package main

import (
	"github.com/qmstr/synclib/pkg/module/rabbitmq"
	"log"
	"os"
	"time"
)

func main() {

	// Connecting to RabbitMQ
	conn, ch := rabbitmq.Connect(os.Getenv("RABBITMQ_ADDRESS"))
	defer conn.Close()
	defer ch.Close()

	// Declaring queues
	queue, _ := rabbitmq.DeclareQueues(ch, os.Getenv("QUEUE_NAME"))

	// Dummy callback function
	dummyCallbackFunction := func() {
		dummyComputationTimeInSeconds := 15
		log.Printf("Performing dummy computation for %d seconds...", dummyComputationTimeInSeconds)
		time.Sleep(time.Duration(dummyComputationTimeInSeconds) * time.Second)
	}

	// Executing the callback function upon receiving the order message
	responded := rabbitmq.OnMessageReceive(ch, queue.Name, dummyCallbackFunction)
	<-responded
}
