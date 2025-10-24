package db

import (
	"fmt"
	"log"
	"os"

	"qr-code-generator/internal/models" //models

	"github.com/joho/godotenv" //environment variables
	"gorm.io/driver/mysql"     // mysql
	"gorm.io/gorm"             // ORM
)

var DB *gorm.DB

// Intitalize MYSQL Connection
func InitDB() {
	//load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back")
	}

	// build DSN
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate schema
	if err := DB.AutoMigrate(&models.QRRecord{}); err != nil {
		log.Fatalf("Failed to migrate schema : %v", err)
	}

	log.Println("Connected to MYSQL & migrated schema")
}
