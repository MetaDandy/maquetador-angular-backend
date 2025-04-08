package config

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
	//&models.User{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}
}
