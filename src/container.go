package src

import (
	"github.com/MetaDandy/maquetador-angular-backend/config"
	"github.com/MetaDandy/maquetador-angular-backend/src/handlers"
	"github.com/MetaDandy/maquetador-angular-backend/src/modules/user"
	"github.com/MetaDandy/maquetador-angular-backend/src/repository"
	"github.com/MetaDandy/maquetador-angular-backend/src/services"
)

type Container struct {
	// Users
	UserRepo    *user.Repository
	UserService *user.Service
	UserHandler *user.Handler

	// Project
	ProjectRepo    *repository.ProjectRepositories
	ProjectService *services.ProjectService
	ProjectHandler *handlers.ProjectHandler
}

func SetupContainer() *Container {
	// User
	userRepo := user.NewRepository(config.DB)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	// Project
	projectRepo := repository.NewProjectRepositories(config.DB)
	projectService := services.NewProjectService(projectRepo, userService)
	projectHandler := handlers.NewProjectHandler(projectService)

	return &Container{
		UserRepo:       userRepo,
		UserService:    userService,
		UserHandler:    userHandler,
		ProjectRepo:    projectRepo,
		ProjectService: projectService,
		ProjectHandler: projectHandler,
	}
}
