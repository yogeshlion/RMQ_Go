package rmq
import (
	"fmt"
	"time"
	"log"
	"github.com/streadway/amqp"
)
func SendConn(body string,cha string){
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch1,_ := conn.Channel()
	q1, _ := ch1.QueueDeclare(cha,false,false,false,false,nil)
	for i:=0;i<2;i++{
		err:= ch1.Publish("",q1.Name,false,false,amqp.Publishing{ContentType: "text/plain",Body: []byte(body)})
		if err!=nil{
			fmt.Println(err)
		}
		time.Sleep(1000*time.Millisecond)
		fmt.Println("Sent: Data")
	}
}
func RevConn(cha string){
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch1,_ := conn.Channel()
	q1, _ := ch1.QueueDeclare(cha,false,false,false,false,nil)
	msgs, _ := ch1.Consume(q1.Name,"",true,false,false,false,nil)
	go func() {
		for d := range msgs {
			time.Sleep(500*time.Millisecond)
			log.Printf("Received: %s", d.Body)
		}
	}()
}
