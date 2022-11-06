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
	publish := rabbit.PublishConfig{
		Exchange:    "school_space",
		RoutingKey:  "",
		Mandatory:   false,
		Immediate:   false,
		ContentType: "application/json",
	}
	rabbitMQ := rabbit.MQConfig{
		Host:     "localhost",
		Port:     "5672",
		User:     "guest",
		Password: "guest",
	}

	publisher := rabbit.NewPublisher(&rabbitMQ)
	publisher.SetupExchange(&exchange)
	publisher.Publish(&publish, []byte(`{id : "1234567810", name: "Prince of Wales", phone: "+23276293366", address: "King Town"}`))
}
