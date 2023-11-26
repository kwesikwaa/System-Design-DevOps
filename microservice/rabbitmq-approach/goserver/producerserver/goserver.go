package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	Userid      string  `json:"userid"`
	Orderid     string  `json:"orderid"`
	Productname string  `json:"productname"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Total       float64 `json:"total"`
}

func main() {

	rest := gin.Default()

	rest.POST("/order", orderHandler)

	err := rest.Run(":8000")
	if err != nil {
		// log error message
	}
}

func orderHandler(c *gin.Context) {
	// message := c.Param("message")
	//assumin the message came through as bellow
	message := Order{
		Userid:      "jlkj3332k",
		Orderid:     "dsfklsjf113",
		Productname: "toy car",
		Quantity:    4,
		Price:       25.00,
		Total:       100.00,
	}

	// order := Order
	// if err := c.ShouldBindJSON(&order); err != nil{
	// 	c.JSON(http.StatusBadRequest, gin.H({"error":err.Error()}))
	// 	return
	// }

	if err := queueHandler(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit Order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order successfully placed"})
}

func queueHandler(order Order) error {
	//create connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Connected to Rabbit..")

	defer conn.Close()

	//create channel
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	//declare a queue
	q, err := ch.QueueDeclare("Order_Queue", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	// convert
	corder, err := json.Marshal(order)
	if err != nil {
		// log
	}
	//publish to the queue
	err = ch.Publish(
		"", q.Name, false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        corder,
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully published to queeu")

	return nil
}
