package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "user=postgres password=VibeSocial42! host=db.axbxrkgbggekcrraugfk.supabase.co port=5432 dbname=postgres"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Event{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = database
}
