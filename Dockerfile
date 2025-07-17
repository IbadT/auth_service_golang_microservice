FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth_service ./cmd/server/main.go

FROM debian:latest

# Устанавливаем зависимости (например, для работы с PostgreSQL)
RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /root/

# Копируем бинарник из builder-этапа
COPY --from=builder /app/auth_service .

# Указываем переменные окружения (можно взять из `docker-compose.yml`)
ENV DB_HOST=auth_db
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV DB_NAME=auth_mic
ENV DB_PORT=5432

# Запуск приложения
CMD ["./auth_service"]
