
package main

import (
	"fmt"
	"consul"
	"erro"
	"rmq"
)
func consulHandle()(consul.Client1,string){
	var addr string
	addr= "127.0.0.1:8500"
	_,inter2,_:=consul.NewConsulClient(addr)
	err1:=inter2.Register("app2",5357)
	fmt.Println("Connected to Consul server....")
	erro.ErrHandle(err1)
	return inter2,"app2"
}

func faultHandle(inter2 consul.Client1,s string,x chan int){
	 inter2.Service(s,x)
}
func main() {
	x:=make(chan int)
	inter2,s:=consulHandle()
	go faultHandle(inter2,s,x)
	go rmq.SendConn("Data","Hello")
	<-x
	fmt.Println("Consul Connection is down...Quitting...")
}
