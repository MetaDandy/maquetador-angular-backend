package project

import (
	"time"

	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"github.com/MetaDandy/maquetador-angular-backend/src/modules/user"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type ProjectCreate struct {
	Title       string         `json:"title" validate:"required"`
	Description string         `json:"description"`
	Content     datatypes.JSON `json:"content"`
	OwnerID     uuid.UUID      `json:"owner_id" validate:"required"`
}

type ProjectUpdate struct {
	Title       *string         `json:"title"`
	Description *string         `json:"description"`
	Content     *datatypes.JSON `json:"content"`
}

type ProjectResponse struct {
	ID          string            `json:"id"`
	Title       string            `json:"title" validate:"required"`
	Description string            `json:"description"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	DeletedAt   *time.Time        `json:"deleted_at"`
	Owner       user.UserResponse `json:"owner"`
}

func ProjectToDTO(u *models.Project) ProjectResponse {
	var deletedAt *time.Time
	if u.DeletedAt.Valid {
		t := u.DeletedAt.Time
		deletedAt = &t
	}

	return ProjectResponse{
		ID:          u.ID.String(),
		Title:       u.Title,
		Description: u.Description,
		Owner:       user.UserToDTO(&u.Owner),
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}

func ProjectsToListDTO(list []models.Project) []ProjectResponse {
	out := make([]ProjectResponse, len(list))
	for i := range list {
		out[i] = ProjectToDTO(&list[i])
	}
	return out
}
