package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "postgres://user:password@postgres:5432/users?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Ошибка соединения с БД: %v", err)
	}

	log.Println("Успешное подключение к PostgreSQL!")
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Ошибка закрытия соединения с БД: %v", err)
	}
}
