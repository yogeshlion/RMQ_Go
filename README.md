# RMQ_Go
Aim:
The aim of this branch is to achieve implementation of a network with RabbitMQ as message broker and Consul as a Service Discovery and Health Check entity.

Used:
1. Go- Version:1.15(used). 
2. RMQ(v3.8.9)- with Erlang support(v22.3)
3. HashiCorp's Consul client(v1.8.4)

File Behavior:
1. applications(sapp,rapp): Any applications in place of these with necessary imports
2. rmq-setup: Setup and functional file for rmq functions of sending receiving and buffering data and control messages
3. consul-setup: Setup and functional file for Consul initiation, Service Discovery and periodic Health Checks
4. error-deal: Error Handling based on requirements of application and use case

Goals:
1. Achieve buffered communication link between 2 applications using RMQ support
2. Concurrency in message sending and reception
3. Dial-up connect/Registry to a Consul Server(Local)
4. Checking Application Status on Consul
5. Periodic HealthChecks by nodes of the consul cluster
6. Other nodes intimated on consul connection failure of a node

External Packages:
1. streadway/amqp
2. hashicorp/consul


