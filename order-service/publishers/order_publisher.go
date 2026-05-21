package publishers

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishOrderCreated(order interface{}) {
	conn, err := amqp.Dial(
		"amqp://guest:guest@rabbitmq:5672/",
	)

	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	channel, err := conn.Channel()

	if err != nil {
		log.Println(err)
		return
	}

	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"order.created",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println(err)
		return
	}

	body, err := json.Marshal(order)

	if err != nil {
		log.Println(err)
		return
	}

	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(
		"Order created event published",
	)
}
