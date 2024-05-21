package utils

import (
	"fmt"
	"github.com/wagslane/go-rabbitmq"
)

func SetupRabbitMQ() (*rabbitmq.Publisher, *rabbitmq.Conn, error) {
	// Load RabbitMQ connection details from environment variables or configuration
	// Create a new RabbitMQ connection
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create RabbitMQ connection: %w", err)
	}

	// Create a new RabbitMQ publisher
	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("emails"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		conn.Close() // Attempt to close the connection if we cannot create the publisher
		return nil, nil, fmt.Errorf("failed to create RabbitMQ publisher: %w", err)
	}

	return publisher, conn, nil
}
