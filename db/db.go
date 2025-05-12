package db

import (
	"devnotes/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=admin password=admin dbname=devnotes_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("🚀 Connected to PostgreSQL database")

	err = DB.AutoMigrate(&models.Note{})
	if err != nil {
		log.Fatal(" Failed to migrate database:", err)
	}
}
