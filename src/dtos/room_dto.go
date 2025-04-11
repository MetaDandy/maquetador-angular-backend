package dtos

import (
	"time"

	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"github.com/google/uuid"
)

type RoomCreate struct {
	ProjectID uuid.UUID `json:"project_id"`
	RoomCode  string    `json:"room_code"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type RoomResponse struct {
	RoomCode  string    `json:"room_code"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`

	Project models.Project
}
