package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title        string    `gorm:"uniqueIndex" json:"title"`
	Description  string    `json:"description"`
	IsPublicLink bool      `gorm:"default:false" json:"is_public_link"`
	PublicToken  string    `gorm:"uniqueIndex" json:"public_token"`
	OwnerID      uuid.UUID `json:"owner_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Owner         User                  `gorm:"foreignKey:OwnerID" json:"owner"`
	Collaborators []ProjectCollaborator `gorm:"foreignKey:ProjectID"`
	Invites       []ProjectInvite       `gorm:"foreignKey:ProjectID"`
	Rooms         []Room                `gorm:"foreignKey:ProjectID"`
}
