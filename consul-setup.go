package consul
import (
	"fmt"
	"time"
	consul "github.com/hashicorp/consul/api"
)
type Client interface {
	Service(string, string) ([]string, error)
	Register(string, int) error
	DeRegister(string) error
}
type Client1 struct {
	consul *consul.Client
}

type Node struct {
    ID              string
    Node            string
    Address         string
    Datacenter      string
    TaggedAddresses map[string]string
    Meta            map[string]string
    CreateIndex     uint64
    ModifyIndex     uint64
}
type HealthCheck struct {
    Node        string
    CheckID     string
    Name        string
    Status      string
    Notes       string
    Output      string
    ServiceID   string
    ServiceName string
    ServiceTags []string
    Type        string
}
type AgentService struct {
    ID                string
    Service           string
    Tags              []string
    Meta              map[string]string
    Port              int
    Address           string
    EnableTagOverride bool
}

type ServiceEntry struct {
    Node    *Node
    Service *AgentService
    Checks  HealthCheck
}
func errHandle(err error){
	if err!=nil{
		fmt.Println("Error Encountered.")
	}
}
//NewConsul returns a Client interface for given consul address
func NewConsulClient(addr string)(*consul.Client, Client1,error) {
	config := consul.DefaultConfig()
	config.Address = addr
	c,_:= consul.NewClient(config)
	var F Client1
	F.consul=c
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

func (c *Client1) Service(service string,x chan int,cnt int){
	time.Sleep(1000*time.Millisecond)
	var cntt int
	for{
		a,_,_:=c.consul.Agent().AgentHealthServiceByName(service)
		fmt.Println("Health Status:",a)
		cntt=cntt+1
		time.Sleep(2000*time.Millisecond)
		if cnt==cntt{
			err:=c.DeRegister(service)
			errHandle(err)
		}
		if a!="passing"{
			x<-cnt
			break
		}
	}
}