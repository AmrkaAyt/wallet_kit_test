package mq

import (
	"encoding/json"
	"log"
	"wallet_kit_test/OrderService/db"

	"github.com/streadway/amqp"
)

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func StartConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Ошибка подключения к RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Ошибка при открытии канала: %v", err)
	}
	defer ch.Close()

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
		log.Fatalf("Ошибка подписки на очередь: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var user User
			err := json.Unmarshal(d.Body, &user)
			if err != nil {
				log.Printf("Ошибка при разборе сообщения: %v", err)
				continue
			}

			err = db.SaveUser(db.User(user))
			if err != nil {
				log.Printf("Ошибка при сохранении пользователя в базу данных: %v", err)
			} else {
				log.Printf("Пользователь %s успешно сохранен в базу данных", user.UserID)
			}
		}
	}()

	log.Printf("Consumer запущен, ожидание сообщений...")
	<-forever
}
