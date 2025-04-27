package user

import (
	"fmt"
	"log"

	"github.com/MetaDandy/maquetador-angular-backend/helper"
	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(req UserRequest) (*UserResponse, *LoginResponse, error) {
	existingUser, err := s.repo.FindByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Error al buscar usuario:", err)
		return nil, nil, err
	}
	if existingUser != nil {
		log.Printf("Correo electrónico en uso: %v\n", existingUser)
		return nil, nil, fmt.Errorf("Correo electrónico en uso")
	}

	id := uuid.New()
	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		log.Printf("Error al hashear la contraseña: %v\n", err)
		return nil, nil, err
	}

	user := &models.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.repo.Create(user); err != nil {
		log.Printf("Error al crear el usuario en la base de datos: %v\n", err)
		return nil, nil, err
	}

	token, err := helper.GenerateJwt(user.ID.String(), string(user.Email))
	if err != nil {
		log.Printf("Error al generar el token JWT: %v\n", err)
		return nil, nil, err
	}

	dto := UserToDTO(user)
	return &dto, &LoginResponse{
		Token: token,
	}, nil
}

func (s *Service) Login(req LoginRequest) (*LoginResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if !helper.CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("Invalid Credentials")
	}

	token, err := helper.GenerateJwt(user.ID.String(), string(user.Email))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: token,
	}, nil
}

func (s *Service) FindUserById(id string) (*UserResponse, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	dto := UserToDTO(user)
	return &dto, nil
}

func (s *Service) FindByEmail(email string) (*UserResponse, error) {
	user, err := s.repo.FindById(email)
	if err != nil {
		return nil, err
	}

	dto := UserToDTO(user)
	return &dto, nil
}

func (s *Service) GetAllUsers(opts *helper.FindAllDto) (*helper.PaginatedResponse[UserResponse], error) {
	users, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := UsersToListDTO(users)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[UserResponse]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) Delete(id string) (*UserResponse, error) {
	user, err := s.FindUserById(id)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}

	return user, nil
}
