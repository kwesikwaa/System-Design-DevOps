package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main(){
	fmt.Println("init...")

	conn,err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil{
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Connected to Rabbit..")

	defer conn.Close()

	ch,err := conn.Channel()
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q,err :=  ch.QueueDeclare("FirstQueue",false,false,false,false,nil)
	if err != nil{
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"","FirstQueue",false,false,ampq.Publishing{
			ContentType: "text/plain",
			Body: []byte("Chale Im on your queue")
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully published")

}