
package main

import (
	"log"
	"fmt"
	"consul"
	"time"
	"github.com/streadway/amqp"
)
func consulHandle()(consul.Client1,string){
	var addr string
	addr= "127.0.0.1:8500"
	_,inter2,_:=consul.NewConsulClient(addr)
	err1:=inter2.Register("app2",5357)
	fmt.Println("Connected to Consul server....")
	fmt.Println(err1)
	return inter2,"app2"
}
func faultHandle(inter2 consul.Client1,s string){
	 inter2.Service(s)
}
func rmqConn() (amqp.Queue,*amqp.Channel){
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch1,_ := conn.Channel()
	q1, _ := ch1.QueueDeclare("hello",false,false,false,false,nil)
	return q1,ch1
}
func application1(body string,q1 amqp.Queue,ch1 *amqp.Channel){
	for i:=0;i<10;i++{
		err := ch1.Publish("",q1.Name,false,false,amqp.Publishing{ContentType: "text/plain",Body: []byte(body)})
		time.Sleep(1000*time.Millisecond)
		fmt.Println(err)
		log.Printf(" [x] Sent %s", body)
}}
func main() {
	inter2,s:=consulHandle()
	q,ch:= rmqConn()
	go faultHandle(inter2,s)
	go application1("Data",q,ch)
	for{
	}
}
