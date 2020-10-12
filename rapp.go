package main

import (
	"fmt"
	"log"
	"consul"
	"time"
	"github.com/streadway/amqp"
)
func consulHandle()(consul.Client1,string){
	var addr string
	addr= "127.0.0.1:8500"
	_,inter2,_:=consul.NewConsulClient(addr)
	err1:=inter2.Register("app1",8000)
	fmt.Println("Connected to Consul server....")
	fmt.Println(err1)
	return inter2,"app1"
}
func faultHandle(inter2 consul.Client1,s string){
	 inter2.Service(s)
}
func rmqConn()(amqp.Queue,*amqp.Channel){
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch1,_ := conn.Channel()
	q1, _ := ch1.QueueDeclare("hello",false,false,false,false,nil)
	return q1,ch1
	}
func application2(q1 amqp.Queue,ch1 *amqp.Channel){
	msgs, _ := ch1.Consume(q1.Name,"",true,false,false,false,nil)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			time.Sleep(500*time.Millisecond)
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf("Waiting for messages....")
	<-forever
}
func main() {
	inter2,s:=consulHandle()
	q,ch:=rmqConn()
	go faultHandle(inter2,s)
	go application2(q,ch)
	for{
	}
}
	