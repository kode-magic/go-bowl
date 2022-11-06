package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func (p *Publisher) Consumers(cfg *ConsumerConfig, queue *amqp.Queue) {
	log.Printf("Declaring consumer: %s", cfg.Consumer)
	ch, chErr := p.Ampq.Channel()
	if chErr != nil {
		panic(chErr)
	}

	deliveries, err := ch.Consume(
		queue.Name,    // queue
		cfg.Consumer,  // consumer
		cfg.AtoAck,    // auto-ack default to true
		cfg.Exclusive, // exclusive
		cfg.NoLocal,   // no-local
		cfg.NoWait,    // no-wait
		nil,           // args
	)

	FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range deliveries {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
