package database

import (
	"log"

	"fiber_go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	dsn := "user=username dbname=catfoodstore sslmode=disable password=yourpassword"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}

	DB = db

	// Автоматическая миграция таблиц
	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Printf("Migration error: %v", err)
	}

	log.Println("Успешно подключились к базе данных")
	return nil
}
