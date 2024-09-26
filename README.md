# Wallet Kit Microservices

Этот проект представляет собой набор микросервисов для управления пользователями и заказами, взаимодействующих через RabbitMQ и gRPC.

## Микросервисы

1. **UserService** – сервис для управления пользователями.
2. **OrderService** – сервис для управления заказами.

Микросервисы взаимодействуют между собой через gRPC и RabbitMQ.

## Структура проекта

- `UserService/` – код микросервиса для управления пользователями.
- `OrderService/` – код микросервиса для управления заказами.
- `docker-compose.yml` – файл для запуска всей инфраструктуры, включая базы данных, очереди сообщений и оба микросервиса.

## Требования

Для запуска проекта необходимы:

- **Docker** и **Docker Compose**
- **Go 1.21**
- **Git** (для локальной разработки)

## Настройка и запуск

1. **Клонируйте репозиторий**:
   ```bash
   git clone https://github.com/AmrkaAyt/wallet_kit_test.git
   cd wallet_kit_test
2. **Запустите микросервисы и необходимое окружение**:
   ```bash
   docker-compose up --build -d
2. **Запустите микросервисы и необходимое окружение**:
   ```bash
   docker-compose up --build -d
   ```

3. **Проверьте статус микросервисов**:
   ```bash
   docker-compose ps
   ```

4. **Просмотрите логи для отладки**:
   ```bash
   docker-compose logs user_service
   docker-compose logs order_service
   ```

## Тестирование микросервисов

### UserService (gRPC запросы)

1. **Создать пользователя**:
   ```bash
   grpcurl -plaintext -d '{"username": "john", "email": "john@example.com"}' localhost:50051 user.UserService/CreateUser
   ```

2. **Получить пользователя по ID**:
   ```bash
   grpcurl -plaintext -d '{"user_id": "UUID_Пользователя"}' localhost:50051 user.UserService/GetUser
   ```

### OrderService (gRPC запросы)

1. **Создать заказ**:
   ```bash
   grpcurl -plaintext -d '{"user_id": "UUID_Пользователя", "product_id": "123", "quantity": 2}' localhost:50052 order.OrderService/CreateOrder
   ```

2. **Получить заказ по ID**:
   ```bash
   grpcurl -plaintext -d '{"order_id": "ORDER_ID"}' localhost:50052 order.OrderService/GetOrder
   ```

## RabbitMQ

Для проверки очереди сообщений можно использовать RabbitMQ Management UI:
- URL: [http://localhost:15672](http://localhost:15672)
- Логин: `guest`
- Пароль: `guest`

## Остановка микросервисов

Чтобы остановить микросервисы и окружение, выполните:
```bash
docker-compose down
```

## Структура баз данных

- **PostgreSQL** – для хранения данных пользователей (UserService).
- **MongoDB** – для хранения заказов (OrderService).

## Автор

Amirzhan Aytnalin(AmrkaAyt)
```

Этот текст завершает файл `README.md`, предоставляя инструкции по установке, тестированию и управлению микросервисами.
