package repository

import (
	"github.com/MetaDandy/maquetador-angular-backend/pkg"
	"github.com/MetaDandy/maquetador-angular-backend/src/dtos"
	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectRepositories struct {
	db *gorm.DB
}

func NewProjectRepositories(db *gorm.DB) *ProjectRepositories {
	return &ProjectRepositories{db}
}

func (r *ProjectRepositories) CreateRepository(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *ProjectRepositories) FindByTitles(title string) (*models.Project, error) {
	var project models.Project
	err := r.db.First(&project, "title = ?", title).Error
	return &project, err
}

func (c *ProjectRepositories) FindAll(opts *dtos.FindAllDto) ([]models.Project, int64, error) {
	var project []models.Project
	query := c.db.Model(models.Project{})
	var total int64
	query, total = pkg.ApplyFindAllOptions(query, opts)

	err := query.Find(&project).Error
	return project, total, err
}

func (r *ProjectRepositories) UpdateRepostory(project *models.Project) error {
	return r.db.Save(project).Error
}

func (r *ProjectRepositories) DeleteRepositories(id uuid.UUID) error {
	return r.db.Delete(&models.Project{}, "id = ?", id).Error
}
