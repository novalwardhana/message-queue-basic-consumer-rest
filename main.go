package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Message Queue Basic Consumer Rest")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		fmt.Printf("An error occured: %s\n", err.Error())
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("An error occured: %s\n", err.Error())
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestNovalQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("An error occured: %s\n", err.Error())
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Println("Receive message: ", string(msg.Body))
		}
	}()

	<-forever
}
