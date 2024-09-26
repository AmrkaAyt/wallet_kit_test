package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"github.com/AmrkaAyt/wallet_kit_test/OrderService/db"
	"github.com/AmrkaAyt/wallet_kit_test/OrderService/internal/service"
	"github.com/AmrkaAyt/wallet_kit_test/OrderService/mq"
	"github.com/AmrkaAyt/wallet_kit_test/OrderService/proto"
	"google.golang.org/grpc"
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

	// Включаем поддержку рефлексии
	reflection.Register(s)

	go mq.StartConsumer()

	log.Println("OrderService is running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
