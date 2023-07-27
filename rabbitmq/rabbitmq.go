package rabbitmq

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
	"notification/token"
)

const (
	QueueName = "notification_queue"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Receive(callback func(*token.Record)) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			record, err := token.DecodeToken(string(d.Body))
			if err != nil {
				log.Printf("Failed to decode token: %s", err)
			} else {
				fmt.Printf("Decoded record: %+v\n", record)
				callback(record)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages.")
	select {}
}
