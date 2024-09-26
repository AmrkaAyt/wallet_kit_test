package mq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func InitRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return nil, nil, fmt.Errorf("Ошибка подключения к RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, fmt.Errorf("Ошибка при открытии канала: %v", err)
	}

	// Объявление очереди для отправки сообщений
	_, err = ch.QueueDeclare(
		"user_created", // имя очереди
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // аргументы
	)
	if err != nil {
		conn.Close()
		ch.Close()
		return nil, nil, fmt.Errorf("Ошибка при объявлении очереди: %v", err)
	}

	return conn, ch, nil
}

func PublishUserCreatedMessage(userID, username, email string) error {
	conn, ch, err := InitRabbitMQ()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()

	user := User{
		UserID:   userID,
		Username: username,
		Email:    email,
	}

	body, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("Ошибка при кодировании данных пользователя в JSON: %v", err)
	}

	// Отправка сообщения
	err = ch.Publish(
		"",             // exchange
		"user_created", // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("Ошибка при отправке сообщения: %v", err)
	}

	log.Printf("Сообщение о создании пользователя %s отправлено в RabbitMQ", userID)
	return nil
}
