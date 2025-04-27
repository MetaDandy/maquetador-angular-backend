package project

import (
	"github.com/MetaDandy/maquetador-angular-backend/helper"
	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *Repository) FindByID(id string) (*models.Project, error) {
	var project models.Project
	err := r.db.Preload("Owner").First(&project, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &project, err
}

func (r *Repository) FindByTitles(title string) (*models.Project, error) {
	var project models.Project
	err := r.db.Preload("Owner").First(&project, "title = ?", title).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &project, err
}

func (c *Repository) FindAll(opts *helper.FindAllDto) ([]models.Project, int64, error) {
	var project []models.Project
	query := c.db.Preload("Owner").Model(models.Project{})
	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&project).Error
	return project, total, err
}

func (r *Repository) FindAllByUserID(ownerID string, opts *helper.FindAllDto) ([]models.Project, int64, error) {
	var projects []models.Project
	query := r.db.Preload("Owner").Where("owner_id = ?", ownerID).Model(&models.Project{})
	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Preload("Owner").Find(&projects).Error
	return projects, total, err
}

func (r *Repository) Update(project *models.Project) error {
	return r.db.Save(project).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&models.Project{}, "id = ?", id).Error
}
