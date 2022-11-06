package rabbit

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

// SetupExchange create message exchange
func (p *Publisher) SetupExchange(cfg *ExchangeConfig) {
	log.Printf("Declaring exchange: %s", cfg.Exchange)

	ch, chErr := p.Ampq.Channel()
	if chErr != nil {
		panic(chErr)
	}
	exErr := ch.ExchangeDeclare(
		cfg.Exchange,
		cfg.Kind,
		cfg.Durable, // default true
		cfg.AutoDelete,
		cfg.Internal,
		cfg.NoWait,
		nil,
	)

	FailOnError(exErr, "Failed to open a channel")
}

func (p *Publisher) SetupQueue(cfg *QueueConfig) *amqp.Queue {
	log.Printf("Declaring queue: %s", cfg.QueueName)

	ch, chErr := p.Ampq.Channel()
	if chErr != nil {
		fmt.Println("chan error: ", chErr)
		panic(chErr)
	}

	q, err := ch.QueueDeclare(
		cfg.QueueName,  // name default ""
		cfg.Durable,    // durable default false
		cfg.AutoDelete, // delete when unused default false
		cfg.Exclusive,  // exclusive default true
		cfg.NoWait,     // no-wait default false
		nil,            // arguments
	)

	FailOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,         // queue name
		cfg.RoutingKey, // routing key default ""
		cfg.Exchange,   // exchange
		cfg.NoWait,
		nil,
	)

	FailOnError(err, "Failed to bind a queue")

	return &q
}
