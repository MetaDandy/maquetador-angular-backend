package project

import (
	"github.com/MetaDandy/maquetador-angular-backend/helper"
	"github.com/MetaDandy/maquetador-angular-backend/src/models"
	"github.com/MetaDandy/maquetador-angular-backend/src/modules/user"
	"gorm.io/gorm"
)

type Service struct {
	repo     *Repository
	userRepo *user.Repository
}

func NewService(repo *Repository, userRepo *user.Repository) *Service {
	return &Service{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *Service) CreateProject(input ProjectCreate) (*ProjectResponse, error) {
	user, err := s.userRepo.FindById(input.OwnerID.String())
	if err != nil {
		return nil, err
	}

	var project *models.Project
	err = s.repo.db.Transaction(func(tx *gorm.DB) error {
		project = &models.Project{
			Title:       input.Title,
			Description: input.Description,
			Content:     input.Content,
			OwnerID:     user.ID,
		}

		if err := tx.Create(project).Error; err != nil {
			return err
		}

		if err := tx.Model(&project).Association("Owner").Append(user); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	reloaded, _ := s.repo.FindByID(project.ID.String())
	dto := ProjectToDTO(reloaded)
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

func (s *Service) UpdateProject(id string, input ProjectUpdate) (*ProjectResponse, error) {
	project, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if input.Title != nil {
		project.Title = *input.Title
	}
	if input.Description != nil {
		project.Description = *input.Description
	}
	if input.Content != nil {
		project.Content = *input.Content
	}

	err = s.repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&project).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	dto := ProjectToDTO(project)
	return &dto, nil
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
