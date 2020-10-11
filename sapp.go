
package main

import (
	"log"
	"fmt"
	"reflect"
	"consul"
	"github.com/streadway/amqp"
)
func consulHandle(){
	var addr string
	addr= "127.0.0.1:8500"
	inter1,inter2,_:=consul.NewConsulClient(addr)
	fmt.Println(inter1)
	fmt.Println(inter2)
	fmt.Println(reflect.TypeOf(inter2))
	fmt.Println(reflect.TypeOf(inter1))
	err1:=inter2.Register("app2",5357)
	fmt.Println(err1)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
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

	body := "Hi"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
