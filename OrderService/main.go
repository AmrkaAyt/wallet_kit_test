package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"wallet_kit_test/OrderService/db"
	"wallet_kit_test/OrderService/internal/service"
	"wallet_kit_test/OrderService/mq"
	"wallet_kit_test/OrderService/proto"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterOrderServiceServer(s, &service.OrderService{})

	go mq.StartConsumer()

	log.Println("OrderService is running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
