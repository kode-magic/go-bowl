package rabbit

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConn(cfg *MQConfig) (con *amqp.Connection, err error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)
	return amqp.Dial(connAddr)
}
