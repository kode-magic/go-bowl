package main

import "github.com/kode-magic/go-bowl/rabbit"

func main() {
	exchange := rabbit.ExchangeConfig{
		Exchange:   "school_space",
		Kind:       "fanout",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
	}
	queue := rabbit.QueueConfig{
		Exchange:   "school_space",
		RoutingKey: "",
		QueueName:  "",
		Durable:    false,
		AutoDelete: false,
		Exclusive:  true,
		NoWait:     false,
	}
	rabbitMQ := rabbit.MQConfig{
		Host:     "localhost",
		Port:     "5672",
		User:     "guest",
		Password: "guest",
	}
	consumer := rabbit.ConsumerConfig{
		Consumer:  "account",
		AtoAck:    true,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Args:      nil,
	}

	publisher := rabbit.NewPublisher(&rabbitMQ)
	publisher.SetupExchange(&exchange)
	q := publisher.SetupQueue(&queue)
	publisher.Consumers(&consumer, q)
}
