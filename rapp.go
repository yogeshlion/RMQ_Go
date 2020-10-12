package main

import (
	"fmt"
	"erro"
	"consul"
	"rmq"
)
func consulHandle()(consul.Client1,string){
	var addr string
	addr= "127.0.0.1:8500"
	_,inter2,_:=consul.NewConsulClient(addr)
	err1:=inter2.Register("app1",8000)
	fmt.Println("Connected to Consul server.....")
	erro.ErrHandle(err1)
	return inter2,"app1"
}
func faultHandle(inter2 consul.Client1,s string,x chan int){
	 inter2.Service(s,x,8)
}

func main() {
	x:=make(chan int)
	inter2,s:=consulHandle()
	go faultHandle(inter2,s,x)
    go rmq.RevConn("Hello")
	<-x
	fmt.Println("Node Down. Informing Neighbours")
	rmq.SendConn("Control: App1 is down","Hello")
	fmt.Println("Exit-Consul Connection.")
}
	