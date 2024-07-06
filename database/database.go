package database

import (
	"crud-app/entities"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() bool {
	dsn := os.Getenv("DATABASE_CONNECTION_STRING")
	if dsn == "" {
		log.Fatal("DATABASE_CONNECTION_STRING is not set")
		return false
	}
	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return false
	}

	err = DB.AutoMigrate(&entities.Item{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
		return false
	}

	return true
}