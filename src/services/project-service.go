package services

import (
	"github.com/MetaDandy/maquetador-angular-backend/src/dtos"
	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"github.com/MetaDandy/maquetador-angular-backend/src/repository"
	"github.com/google/uuid"
)

type ProjectService struct {
	projectRepo *repository.ProjectRepositories
	userService *UserService
}

func NewProjectService(projectRepo *repository.ProjectRepositories, userService *UserService) *ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
		userService: userService,
	}
}

func (s *ProjectService) CreateProject(req dtos.ProjectCreate) (*dtos.ProjectResponse, error) {
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
		ID:           id,
		Title:        req.Title,
		Description:  req.Description,
		IsPublicLink: req.IsPublicLink,
		PublicToken:  req.PublicToken,
		OwnerID:      ownerID,
		Owner:        userModel,
	}

	if err := s.projectRepo.CreateRepository(project); err != nil {
		return nil, err
	}

	return &dtos.ProjectResponse{
		Title:        project.Title,
		Description:  project.Description,
		Owner:        project.Owner,
		IsPublicLink: project.IsPublicLink,
		PublicToken:  project.PublicToken,
		CreatedAt:    project.CreatedAt,
		UpdatedAt:    project.CreatedAt,
	}, nil
}

func (s *ProjectService) FindByID(id string) (*dtos.ProjectResponse, error) {
	project, err := s.projectRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dtos.ProjectResponse{
		ID:           project.ID,
		Title:        project.Title,
		Description:  project.Description,
		Owner:        project.Owner,
		IsPublicLink: project.IsPublicLink,
		PublicToken:  project.PublicToken,
		CreatedAt:    project.CreatedAt,
		UpdatedAt:    project.UpdatedAt,
	}, nil
}

func (s *ProjectService) FindByTitles(title string) (*dtos.ProjectResponse, error) {
	project, err := s.projectRepo.FindByTitles(title)
	if err != nil {
		return nil, err
	}

	return &dtos.ProjectResponse{
		ID:           project.ID,
		Title:        title,
		Description:  project.Description,
		Owner:        project.Owner,
		IsPublicLink: project.IsPublicLink,
		PublicToken:  project.PublicToken,
		CreatedAt:    project.CreatedAt,
		UpdatedAt:    project.UpdatedAt,
	}, nil
}

func (s *ProjectService) GetAll(opts *dtos.FindAllDto) (*dtos.PaginatedResponse[models.Project], error) {
	data, total, err := s.projectRepo.FindAll(opts)
	if err != nil {
		return nil, err
	}

	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &dtos.PaginatedResponse[models.Project]{
		Data:   data,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *ProjectService) Delete(id string) (*dtos.ProjectResponse, error) {
	project, err := s.FindByID(id)
	if err != nil {
		return nil, err
	}

	if err := s.projectRepo.Delete(id); err != nil {
		return nil, err
	}

	return project, nil
}
