package user

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

func (r *Repository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "email = ?", email).Error
	return &user, err
}

func (r *Repository) FindById(id string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}

func (c *Repository) FindAll(opts *helper.FindAllDto) ([]models.User, int64, error) {
	var users []models.User
	query := c.db.Model(models.User{})
	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&users).Error
	return users, total, err
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}
