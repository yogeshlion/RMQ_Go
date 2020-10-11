package main

import (
	"fmt"
	"log"
	"consul"
	"reflect"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func consulHandle(){
	var addr string
	addr= "127.0.0.1:8500"
	inter1,inter2,_:=consul.NewConsulClient(addr)
	fmt.Println(inter1)
	fmt.Println(inter2)
	fmt.Println(reflect.TypeOf(inter2))
	fmt.Println(reflect.TypeOf(inter1))
	err1:=inter2.Register("app1",8000)
	fmt.Println(err1)
}

func main() {
	consulHandle()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("Waiting for messages....")
	<-forever
}