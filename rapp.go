package main

import (
	"fmt"
	"log"
	"consul"
	//"reflect"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func consulHandle()(consul.Client1,string){
	var addr string
	addr= "127.0.0.1:8500"
	_,inter2,_:=consul.NewConsulClient(addr)
	//fmt.Println(inter1)
	err1:=inter2.Register("app1",8000)
	fmt.Println("Connected to Consul server....")
	//fmt.Println(x)
	fmt.Println(err1)
	return inter2,"app1"
}
func faultHandle(inter2 consul.Client1,s string){
	 inter2.Service(s)
}

func main() {
	inter2,s:=consulHandle()
	faultHandle(inter2,s)
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