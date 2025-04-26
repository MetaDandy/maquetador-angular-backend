package config

import (
	"log"

	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Project{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}
}
