package services

import (
	"fmt"
	"time"

	"github.com/MetaDandy/maquetador-angular-backend/pkg"
	"github.com/MetaDandy/maquetador-angular-backend/src/dtos"
	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"github.com/MetaDandy/maquetador-angular-backend/src/repository"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(req dtos.UserRequest) (*dtos.UserResponse, error) {
	id := uuid.New()
	hashedPassword, err := pkg.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		ID:        id.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) Login(req dtos.LoginRequest) (*dtos.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if !pkg.CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("Invalid Credentials")
	}

	token, err := pkg.GenerateJwt(user.ID.String(), string(user.Email))

	if err != nil {
		return nil, err
	}

	return &dtos.LoginResponse{
		Token: token,
	}, nil

}

func (s *UserService) FindUserById(id string) (*dtos.UserResponse, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) FindByEmail(email string) (*dtos.UserResponse, error) {
	user, err := s.userRepo.FindById(email)
	if err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		ID:        user.Email,
		Name:      user.Name,
		Email:     email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) GetAllUsers(opts *dtos.FindAllDto) (*dtos.PaginatedResponse[models.User], error) {
	data, total, err := s.userRepo.FindAll(opts)
	if err != nil {
		return nil, err
	}

	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &dtos.PaginatedResponse[models.User]{
		Data:   data,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *UserService) Delete(id string) (*dtos.UserResponse, error) {
	user, err := s.FindUserById(id)
	if err != nil {
		return nil, err
	}

	if err := s.userRepo.Delete(id); err != nil {
		return nil, err
	}

	return user, nil
}
