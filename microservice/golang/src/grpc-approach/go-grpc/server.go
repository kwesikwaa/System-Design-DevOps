package main

import (
	"log"
)

const (
	port = ":5000"
)

func errochecker(err error, msg string) {
	log.Fatalf("Error: %v : %v", msg, err)
}

type OrderServer struct {
}

// func (s *OrderServer) CreateOrder(ctx context.Context) error {

// }
