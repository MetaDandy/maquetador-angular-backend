package repository

import (
	"github.com/MetaDandy/maquetador-angular-backend/pkg"
	"github.com/MetaDandy/maquetador-angular-backend/src/dtos"
	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "email = ?", email).Error
	return &user, err
}

func (r *UserRepository) FindById(id string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}

func (c *UserRepository) FindAll(opts *dtos.FindAllDto) ([]models.User, int64, error) {
	var users []models.User
	query := c.db.Model(models.User{})
	var total int64
	query, total = pkg.ApplyFindAllOptions(query, opts)

	err := query.Find(&users).Error
	return users, total, err
}

func (r *UserRepository) Delete(id string) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}
