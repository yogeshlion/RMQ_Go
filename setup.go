package consul
import (
	"fmt"
	"reflect"
	//"time"
	consul "github.com/hashicorp/consul/api"
)
type Client interface {
// Get a Service from consul
	Service(string, string) ([]string, error)
// Register a service with local agent
	Register(string, int) error
// Deregister a service with local agent
	DeRegister(string) error
}

type Client1 struct {
	consul *consul.Client
}
	
//NewConsul returns a Client interface for given consul address
func NewConsulClient(addr string)(*consul.Client, Client1,error) {
	//fmt.Println(reflect.TypeOf(con))
	config := consul.DefaultConfig()
	config.Address = addr
	c,_:= consul.NewClient(config)
	fmt.Println(reflect.TypeOf(c))
	//if err != nil {
	//	return nil,nil, err
	//}
	var F Client1
	fmt.Println(reflect.TypeOf(F))
	F.consul=c
	fmt.Println(F)
	//fmt.Println(&client{consul: c})
	//fmt.Println(reflect.TypeOf(&client{consul: c}))
	return c,F, nil
}
// Register a service with consul local agent
func (c *Client1) Register(name string, port int) error {
	reg := &consul.AgentServiceRegistration{
		ID:   name,
		Name: name,
		Port: port,
	}
	return c.consul.Agent().ServiceRegister(reg)
}

// DeRegister a service with consul local agent
func (c *Client1) DeRegister(id string) error {
	return c.consul.Agent().ServiceDeregister(id)
}