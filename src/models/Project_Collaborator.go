package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectCollaborator struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ProjectID uuid.UUID      `json:"project_id"`
	UserID    uuid.UUID      `json:"user_id"`
	InvitedBy uuid.UUID      `json:"invited_by"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Project Project
	User    User
}
