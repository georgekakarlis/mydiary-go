package database

import (
	"fmt"
	"mydiary/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB

// ConnectDB connects to the database
func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("mydiary.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.User{}, &models.Diary{}, &models.RefreshToken{})
	fmt.Println("Database Migrated")
}