package main

import (
	"log"

	"github.com/IbadT/auth_service_golang_microservice.git/internal/auth"
	"github.com/IbadT/auth_service_golang_microservice.git/internal/database"
	transportgrpc "github.com/IbadT/auth_service_golang_microservice.git/internal/transport"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	repository := auth.NewRepository(db)
	service := auth.NewService(repository)

	if err := transportgrpc.RunGRPC(service); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
