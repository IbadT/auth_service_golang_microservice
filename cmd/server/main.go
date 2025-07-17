package main

import (
	"log"

	"github.com/IbadT/auth_service_golang_microservice/internal/auth"
	"github.com/IbadT/auth_service_golang_microservice/internal/database"
	transportgrpc "github.com/IbadT/auth_service_golang_microservice/internal/transport"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// accessSecret := []byte(os.Getenv("ACCESS_SECRET"))
	// refreshSecret := []byte(os.Getenv("REFRESH_SECRET"))

	// if len(accessSecret) == 0 || len(refreshSecret) == 0 {
	// 	log.Fatal("Secrets are not defined in .env")
	// }

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
