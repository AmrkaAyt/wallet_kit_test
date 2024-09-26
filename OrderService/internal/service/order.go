package service

import (
	"context"
	"fmt"
	"github.com/AmrkaAyt/wallet_kit_test/OrderService/proto"
	userpb "github.com/AmrkaAyt/wallet_kit_test/UserService/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
)

type OrderService struct {
	proto.UnimplementedOrderServiceServer
}

func (s *OrderService) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	conn, err := grpc.Dial("user_service:50051", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к UserService: %v", err)
	}
	defer conn.Close()

	userClient := userpb.NewUserServiceClient(conn)
	userResp, err := userClient.GetUser(ctx, &userpb.GetUserRequest{UserId: req.UserId})
	if err != nil {
		return nil, fmt.Errorf("Ошибка при получении пользователя: %v", err)
	}

	if userResp == nil || userResp.Username == "" {
		return nil, fmt.Errorf("Пользователь с ID %s не найден", req.UserId)
	}

	orderID := uuid.New().String()

	log.Printf("Создан заказ с ID: %s для пользователя: %s", orderID, userResp.Username)

	return &proto.CreateOrderResponse{OrderId: orderID}, nil
}

func (s *OrderService) GetOrder(ctx context.Context, req *proto.GetOrderRequest) (*proto.GetOrderResponse, error) {
	return &proto.GetOrderResponse{
		UserId:    "user-id",
		ProductId: "product-id",
		Quantity:  1,
	}, nil
}
