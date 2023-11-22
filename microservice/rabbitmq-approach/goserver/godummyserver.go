// the idea is to communicate with the python server.
// this is only an illustration of consumer in go

package consumer

import (
	"fmt"
)

func main() {
	fmt.Println("init consumer")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"FirstQueue", "", true, false, false, false, nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("Received Message: %v \n", d.body)
		}
	}()

	fmt.Println("Successfully connected")
	fmt.Println(" ~~ waiting for message")
	<-forever
}
