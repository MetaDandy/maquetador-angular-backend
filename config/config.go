package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	Port string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "3000"
	}

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			log.Fatal("DATABASE_URL not set in .env file")
		}

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Printf("Database connected successfully after %d attempt(s)", i+1)
			Migrate(DB)
			return
		}
		log.Printf("Failed to connect to database, retrying (%d/%d): %v", i+1, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Error connecting to database after %d retries", maxRetries)

	Migrate(DB)
	log.Printf("Database connected")
}
