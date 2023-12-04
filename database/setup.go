package database

import (
	"event-tracking/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := viper.GetString("DATABASE_DSN")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&models.Event{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = database
}
