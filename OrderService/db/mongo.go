package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	UserID   string `bson:"user_id"`  // Поле для хранения ID пользователя
	Username string `bson:"username"` // Поле для хранения имени пользователя
	Email    string `bson:"email"`    // Поле для хранения email пользователя
}

var Client *mongo.Client
var UsersCollection *mongo.Collection

// InitDB - подключение к MongoDB
func InitDB() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatalf("Ошибка подключения к MongoDB: %v", err)
	}

	UsersCollection = Client.Database("orders_db").Collection("users")

	log.Println("Успешное подключение к MongoDB!")
}

func CloseDB() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Ошибка отключения от MongoDB: %v", err)
	}
}

func SaveUser(user User) error {
	_, err := UsersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return fmt.Errorf("Ошибка при сохранении пользователя в MongoDB: %v", err)
	}
	return nil
}
