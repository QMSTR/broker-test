package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func declareQueue(ch *amqp.Channel, queue string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return q
}

func publishMessage(ch *amqp.Channel, q *amqp.Queue, body string) {
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

func main() {

	// Connecting to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq-service:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declaring queues
	builderQueue := declareQueue(ch, "builder")
	analyzerQueue := declareQueue(ch, "analyzer")
	reporterQueue := declareQueue(ch, "reporter")

	type module struct {
		name string
		queue amqp.Queue
	}

	// Modules who are going to be waked up, in order
	modulesToWakeUp := []module{
		{
			name: "builder",
			queue: builderQueue,
		},
		{
			name: "analyzer",
			queue: analyzerQueue,
		},
		{
			name: "reporter",
			queue: reporterQueue,
		},
	}

	// Waking up one Pod of each module at the time
	for _, module := range modulesToWakeUp {
		sleepTimeInSeconds := 5
		log.Printf("Waking up one %s...", module.name)
		publishMessage(ch, &module.queue, fmt.Sprintf("Go, %s!", module.name))
		log.Printf("...%s event sent. Sleeping for %d seconds...",
			module.name, sleepTimeInSeconds)
		time.Sleep(time.Duration(sleepTimeInSeconds) * time.Second)
	}

	log.Printf("...done.")
}
