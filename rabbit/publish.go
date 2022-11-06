package rabbit

import (
	"context"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

// NewPublisher rabbitmq publisher constructor
func NewPublisher(cfg *MQConfig) *Publisher {
	mqConn, connErr := NewRabbitMQConn(cfg)
	FailOnError(connErr, "Fail to connect to RabbitMQ")

	amqpChan, chErr := mqConn.Channel()

	FailOnError(chErr, "Failed to create a channel")

	defer func(amqpChan *amqp.Channel) {
		err := amqpChan.Close()
		if err != nil {
			FailOnError(err, "Failed to create wth open channel")
		}
	}(amqpChan)

	log.Println("Connection initialized")

	return &Publisher{Ampq: mqConn}
}

func (p *Publisher) Publish(cfg *PublishConfig, body []byte) {
	log.Printf("Declaring publisher")
	ch, chErr := p.Ampq.Channel()
	if chErr != nil {
		panic(chErr)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ch.PublishWithContext(ctx,
		cfg.Exchange,   // exchange
		cfg.RoutingKey, // routing key default to ""
		cfg.Mandatory,  // mandatory
		cfg.Immediate,  // immediate
		amqp.Publishing{
			ContentType: cfg.ContentType,
			MessageId:   uuid.New().String(),
			Body:        []byte(body),
		})

	FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
