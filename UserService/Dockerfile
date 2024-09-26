FROM golang:1.21-alpine

# Установка зависимостей
RUN apk add --no-cache git

# Создание рабочей директории
WORKDIR /app

# Копирование файлов go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Установка зависимостей
RUN go mod tidy

# Копирование исходных файлов
COPY . .

# Сборка приложения
RUN go build -o user_service .

# Запуск приложения
CMD ["/app/user_service"]
