package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type Config struct {
	RabbitMQ MQConfig
	Exchange ExchangeConfig
	Queue    QueueConfig
	Publish  PublishConfig
}

// Publisher rabbitmq publisher
type Publisher struct {
	Ampq *amqp.Connection
}

// MQConfig rabbitmq config
type MQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

type PublishConfig struct {
	Exchange    string
	RoutingKey  string
	Mandatory   bool
	Immediate   bool
	ContentType string
}

type ConsumerConfig struct {
	Consumer  string
	AtoAck    bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      interface{}
}

type ExchangeQueueConfig struct {
	eConfig ExchangeConfig
	qConfig QueueConfig
}

type ExchangeConfig struct {
	Exchange   string
	Kind       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
}

type QueueConfig struct {
	Exchange   string
	QueueName  string
	RoutingKey string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
}
