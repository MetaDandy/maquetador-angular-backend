package repository

import (
	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *ProjectRepository) FindByTitle(title string) (*models.Project, error) {
	var project models.Project
	err := r.db.First(&project, "title = ?", title).Error
	return &project, err
}

func (r *ProjectRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Project{}, "id = ?", id).Error
}
