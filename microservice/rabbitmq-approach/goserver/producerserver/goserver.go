package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	errorhandler(err)
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
	errorhandler(err)

	fmt.Println("Connected to Rabbit..")

	defer conn.Close()

	//create channel
	ch, err := conn.Channel()
	errorhandler(err)
	defer ch.Close()

	//declare a queue
	q, err := ch.QueueDeclare("Order_Queue", true, false, false, false, nil)
	errorhandler(err)

	fmt.Println(q)

	// convert
	corder, err := json.Marshal(order)
	errorhandler(err)

	//publish to the queue
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"", q.Name, false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        corder,
		},
	)
	errorhandler(err)
	fmt.Println("Successfully published to queeu")

	return nil
}

func errorhandler(err error) {
	if err != nil {
		// fmt.Println(msg)
		panic(err)
	}
}
