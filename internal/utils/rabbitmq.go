package utils

import (
	"encoding/json"
	"fmt"
	"github.com/wagslane/go-rabbitmq"
	"log"
	rabbitmq2 "madmax/internal/application/db/rabbitmq"
	"madmax/internal/entity"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RabbitConnect() {
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	consumer, err := rabbitmq.NewConsumer(
		conn,
		"my_queue",
		rabbitmq.WithConsumerOptionsRoutingKey("meme.emails"),
		rabbitmq.WithConsumerOptionsExchangeName("emails"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	err = consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
		var mqm entity.MailQueMessage
		err := json.Unmarshal(d.Body, &mqm)
		if err != nil {

		}
		var ok bool
		for {
			ok = rabbitmq2.SendHTML(mqm)
			if ok {
				break
			}
			time.Sleep(time.Second * 2)
		}
		return rabbitmq.Ack
	})
	if err != nil {
		log.Fatal(err)
	}
	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("stopping consumer")

}
