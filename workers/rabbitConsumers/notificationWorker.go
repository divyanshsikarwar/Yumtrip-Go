package rabbitconsumers

import (
	"context"
	"fmt"
	"log"
	"yumtrip/constants"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NotificationConsumer() {
	conn, err := amqp.Dial(constants.RabbitMQURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		constants.RabbitNotificationQueue, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			//Send Email
			fmt.Println(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}


func NotificationProducer (msg string) {
	conn, err := amqp.Dial(constants.RabbitMQURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.
	QueueDeclare(constants.RabbitNotificationQueue, false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	err = ch.PublishWithContext(
		context.Background(),
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "json/Application",
			Body:        []byte(msg),
		})
	failOnError(err, "Failed to publish a message")
}

			