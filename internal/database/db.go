package database

import (
	"log"

	"github.com/IbadT/auth_service_golang_microservice/internal/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=auth_db user=postgres password=postgres dbname=auth_mic port=5432 sslmode=disable"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return nil, err
	}

	if err = DB.AutoMigrate(auth.User{}); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
		return nil, err
	}

	return DB, nil
}
