package models

import (
	"time"

	"github.com/google/uuid"
)

type ProjectInvite struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProjectID uuid.UUID `json:"project_id"`
	Email     string    `json:"email"`
	Token     string    `gorm:"uniqueIndex" json:"token"`
	Status    string    `gorm:"default:'pending'" json:"status"` // 'pending', 'accepted', 'revoked'
	InvitedBy uuid.UUID `json:"invited_by"`
	CreatedAt time.Time `json:"created_at"`

	Project Project
}
