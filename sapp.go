
package main

import (
	"log"
	"fmt"
	//"reflect"
	"consul"
	"github.com/streadway/amqp"
)
func consulHandle()(consul.Client1,string){
	var addr string
	addr= "127.0.0.1:8500"
	_,inter2,_:=consul.NewConsulClient(addr)
	//fmt.Println(inter1)
	err1:=inter2.Register("app2",5357)
	fmt.Println("Connected to Consul server....")
	//fmt.Println(x)
	fmt.Println(err1)
	return inter2,"app2"
}
func faultHandle(inter2 consul.Client1,s string){
	 inter2.Service(s)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	inter2,s:=consulHandle()
	go faultHandle(inter2,s)
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
	for{

	}
}
