package main

import (
	"github.com/AmrkaAyt/wallet_kit_test/UserService/db"
	"github.com/AmrkaAyt/wallet_kit_test/UserService/internal/service"
	pb "github.com/AmrkaAyt/wallet_kit_test/UserService/proto"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	db.InitDB() // Добавьте это, если ещё нет
	defer db.CloseDB()

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &service.UserServiceServer{})

	reflection.Register(s)

	go startRabbitMQ()

	log.Println("UserService is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func startRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"user_created", // имя очереди
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // аргументы
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		"user_created", // имя очереди
		"",             // имя консюмера
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // аргументы
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	forever := make(chan bool)
	<-forever
}
