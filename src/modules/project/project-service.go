package project

import (
	"github.com/MetaDandy/maquetador-angular-backend/helper"
	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"github.com/MetaDandy/maquetador-angular-backend/src/modules/user"
	"github.com/google/uuid"
)

type Service struct {
	repo        *Repository
	userService *user.Service
}

func NewService(repo *Repository, userService *user.Service) *Service {
	return &Service{
		repo:        repo,
		userService: userService,
	}
}

func (s *Service) CreateProject(req ProjectCreate) (*ProjectResponse, error) {
	id := uuid.New()

	user, err := s.userService.FindUserById(req.OwnerID.String())
	if err != nil {
		return nil, err
	}

	ownerID, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, err
	}

	userModel := models.User{
		ID:    ownerID,
		Name:  user.Name,
		Email: user.Email,
	}

	project := &models.Project{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		OwnerID:     ownerID,
		Owner:       userModel,
	}

	if err := s.repo.Create(project); err != nil {
		return nil, err
	}

	dto := ProjectToDTO(project)
	return &dto, nil
}

func (s *Service) FindByID(id string) (*ProjectResponse, error) {
	project, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	dto := ProjectToDTO(project)
	return &dto, nil
}

func (s *Service) FindByTitles(title string) (*ProjectResponse, error) {
	project, err := s.repo.FindByTitles(title)
	if err != nil {
		return nil, err
	}

	dto := ProjectToDTO(project)
	return &dto, nil
}

func (s *Service) GetAll(opts *helper.FindAllDto) (*helper.PaginatedResponse[ProjectResponse], error) {
	data, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := ProjectsToListDTO(data)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[ProjectResponse]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) FindAllProjectsByUser(ownerID string, opts *helper.FindAllDto) (*helper.PaginatedResponse[ProjectResponse], error) {
	data, total, err := s.repo.FindAllByUserID(ownerID, opts)
	if err != nil {
		return nil, err
	}

	dtos := ProjectsToListDTO(data)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[ProjectResponse]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) Delete(id string) (*ProjectResponse, error) {
	project, err := s.FindByID(id)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}

	return project, nil
}
