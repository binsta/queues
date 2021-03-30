package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer Application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("err")
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("err")
		panic(err)
	}
	defer ch.Close()

	msg, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range msg {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()
	fmt.Println("Sucessfully connected to our RabbitMQ instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}
