package dtos

import (
	"time"

	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"github.com/google/uuid"
)

type ProjectCreate struct {
	Title        string    `json:"title" validate:"required"`
	Description  string    `json:"description"`
	OwnerID      uuid.UUID `json:"owner_id" validate:"required"`
	IsPublicLink bool      `json:"is_public_link"`
	PublicToken  string    `json:"public_token" validate:"required"`
}

type ProjectUpdate struct {
	Title        *string    `json:"title"`
	Description  *string    `json:"description"`
	OwnerID      *uuid.UUID `json:"owner_id"`
	IsPublicLink *bool      `json:"is_public_link"`
	PublicToken  *string    `json:"public_token"`
}

type ProjectResponse struct {
	ID           uuid.UUID   `json:"id"`
	Title        string      `json:"title" validate:"required"`
	Description  string      `json:"description"`
	Owner        models.User `json:"owner"`
	IsPublicLink bool        `json:"is_public_link"`
	PublicToken  string      `json:"public_token" validate:"required"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`

	// poner mas datos
}
