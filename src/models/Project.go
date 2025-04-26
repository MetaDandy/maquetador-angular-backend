package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Project struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string         `gorm:"uniqueIndex"`
	Description string         `json:"description"`
	Content     datatypes.JSON `gorm:"type:jsonb"`
	OwnerID     uuid.UUID      `json:"owner_id"`
	CreatedAt   time.Time      `gorm:"not null;default:now()"`
	UpdatedAt   time.Time      `gorm:"not null;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Owner User `gorm:"foreignKey:OwnerID" json:"owner"`
}
