package consul
import (
	"fmt"
	//"reflect"
	"time"
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
type QueryOptions struct {
    // Namespace overrides the `default` namespace
    // Note: Namespaces are available only in Consul Enterprise
    Namespace string

    // Providing a datacenter overwrites the DC provided
    // by the Config
    Datacenter string

    // AllowStale allows any Consul server (non-leader) to service
    // a read. This allows for lower latency and higher throughput
    AllowStale bool

    // RequireConsistent forces the read to be fully consistent.
    // This is more expensive but prevents ever performing a stale
    // read.
    RequireConsistent bool

    // UseCache requests that the agent cache results locally. See
    // https://www.consul.io/api/features/caching.html for more details on the
    // semantics.
    UseCache bool

    // MaxAge limits how old a cached value will be returned if UseCache is true.
    // If there is a cached response that is older than the MaxAge, it is treated
    // as a cache miss and a new fetch invoked. If the fetch fails, the error is
    // returned. Clients that wish to allow for stale results on error can set
    // StaleIfError to a longer duration to change this behavior. It is ignored
    // if the endpoint supports background refresh caching. See
    // https://www.consul.io/api/features/caching.html for more details.
    MaxAge time.Duration

    // StaleIfError specifies how stale the client will accept a cached response
    // if the servers are unavailable to fetch a fresh one. Only makes sense when
    // UseCache is true and MaxAge is set to a lower, non-zero value. It is
    // ignored if the endpoint supports background refresh caching. See
    // https://www.consul.io/api/features/caching.html for more details.
    StaleIfError time.Duration

    // WaitIndex is used to enable a blocking query. Waits
    // until the timeout or the next index is reached
    WaitIndex uint64

    // WaitHash is used by some endpoints instead of WaitIndex to perform blocking
    // on state based on a hash of the response rather than a monotonic index.
    // This is required when the state being blocked on is not stored in Raft, for
    // example agent-local proxy configuration.
    WaitHash string

    // WaitTime is used to bound the duration of a wait.
    // Defaults to that of the Config, but can be overridden.
    WaitTime time.Duration

    // Token is used to provide a per-request ACL token
    // which overrides the agent's default token.
    Token string

    // Near is used to provide a node name that will sort the results
    // in ascending order based on the estimated round trip time from
    // that node. Setting this to "_agent" will use the agent's node
    // for the sort.
    Near string

    // NodeMeta is used to filter results by nodes with the given
    // metadata key/value pairs. Currently, only one key/value pair can
    // be provided for filtering.
    NodeMeta map[string]string

    // RelayFactor is used in keyring operations to cause responses to be
    // relayed back to the sender through N other random nodes. Must be
    // a value from 0 to 5 (inclusive).
    RelayFactor uint8

    // LocalOnly is used in keyring list operation to force the keyring
    // query to only hit local servers (no WAN traffic).
    LocalOnly bool

    // Connect filters prepared query execution to only include Connect-capable
    // services. This currently affects prepared query execution.
    Connect bool

    // Filter requests filtering data prior to it being returned. The string
    // is a go-bexpr compatible expression.
    Filter string
    // contains filtered or unexported fields
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
 //   Namespace   string `json:",omitempty"`
}
type AgentService struct {
    //Kind              ServiceKind `json:",omitempty"`
    ID                string
    Service           string
    Tags              []string
    Meta              map[string]string
    Port              int
    Address           string
   // TaggedAddresses   map[string]ServiceAddress `json:",omitempty"`
 //   Weights           AgentWeights
    EnableTagOverride bool
}

type ServiceEntry struct {
    Node    *Node
    Service *AgentService
    Checks  HealthCheck
}
type QueryMeta struct {
    // LastIndex. This can be used as a WaitIndex to perform
    // a blocking query
    LastIndex uint64

    // LastContentHash. This can be used as a WaitHash to perform a blocking query
    // for endpoints that support hash-based blocking. Endpoints that do not
    // support it will return an empty hash.
    LastContentHash string

    // Time of last contact from the leader for the
    // server servicing the request
    LastContact time.Duration

    // Is there a known leader
    KnownLeader bool

    // How long did the request take
    RequestTime time.Duration

    // Is address translation enabled for HTTP responses on this agent
    AddressTranslationEnabled bool

    // CacheHit is true if the result was served from agent-local cache.
    CacheHit bool

    // CacheAge is set if request was ?cached and indicates how stale the cached
    // response is.
    CacheAge time.Duration
}


	
//NewConsul returns a Client interface for given consul address
func NewConsulClient(addr string)(*consul.Client, Client1,error) {
	//fmt.Println(reflect.TypeOf(con))
	config := consul.DefaultConfig()
	config.Address = addr
	c,_:= consul.NewClient(config)
	//fmt.Println(reflect.TypeOf(c))
	//if err != nil {
	//	return nil,nil, err
	//}
	var F Client1
	//fmt.Println(reflect.TypeOf(F))
	F.consul=c
	//fmt.Println(F)
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


func (c *Client1) Service(service string){
	//passingOnly := true 
	//var addrs1 ServiceEntry
	//q:=QueryOptions{}
	//addrs, meta,_:= c.consul.Health().Service(service, tag, passingOnly, nil)
	//buff:="done"
	time.Sleep(1000*time.Millisecond)
	var cnt int
	for{
		a,_,_:=c.consul.Agent().AgentHealthServiceByName(service)
	//fmt.Println(&addrs)
	//fmt.Println(meta)
		fmt.Println(a)
		cnt=cnt+1
		//fmt.Println(b)
		time.Sleep(2000*time.Millisecond)
		if a!="passing"{
			break
		}
		if cnt==5{
			error:=c.DeRegister(service)
			fmt.Println(error)
		}
	}
	
	//fmt.Println(b)
	//addrs1:=[]*ServiceEntry{}
	//meta1:=[]*QueryMeta{}
	//if len(addrs) == 0 && err == nil {
	//	return nil, fmt.Errorf("service ( %s ) was not found", service)
	//}
	//if err != nil {
	//	return nil, err
	//-}
	
	//fmt.Println(reflect.TypeOf(addrs))
	//fmt.Println(reflect.TypeOf(meta))
	//return 
}