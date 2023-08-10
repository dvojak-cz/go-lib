package messageBroker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQBroker struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (rbmb *RabbitMQBroker) Close() {
	rbmb.ch.Close()
	rbmb.conn.Close()
}

func NewRabbitMQBroker(connectionData IConnectionData) (*RabbitMQBroker, error) {
	conn, err := amqp.Dial(connectionData.ConnectionString())
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQBroker{conn, ch}, nil
}