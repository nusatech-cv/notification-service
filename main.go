package main

import (
	"notification/firebase"
	"notification/rabbitmq"
	"notification/token"
)

func main() {
	rabbitmq.Receive(func(record *token.Record) {
		firebase.SendNotification(record)
	})
}
