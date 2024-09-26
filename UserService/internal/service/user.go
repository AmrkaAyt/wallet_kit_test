package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"wallet_kit_test/UserService/db"
	"wallet_kit_test/UserService/mq"
	pb "wallet_kit_test/UserService/proto"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	userID := uuid.New().String()

	query := `INSERT INTO users (user_id, username, email) VALUES ($1, $2, $3)`
	_, err := db.DB.Exec(query, userID, req.Username, req.Email)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при сохранении пользователя: %v", err)
	}

	err = mq.PublishUserCreatedMessage(userID, req.Username, req.Email)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при отправке сообщения в RabbitMQ: %v", err)
	}

	return &pb.CreateUserResponse{UserId: userID}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var username, email string
	query := `SELECT username, email FROM users WHERE user_id = $1`
	err := db.DB.QueryRow(query, req.UserId).Scan(&username, &email)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при получении пользователя: %v", err)
	}

	return &pb.GetUserResponse{Username: username, Email: email}, nil
}
