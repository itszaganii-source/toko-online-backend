package config

import (
	"log"
	"toko-online-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	
	// Connect to SQLite database
	DB, err = gorm.Open(sqlite.Open("toko_online.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	// Auto migrate the Product model
	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	
	log.Println("Database connected and migrated successfully")
}
