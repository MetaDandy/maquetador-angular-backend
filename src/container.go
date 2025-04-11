package src

import (
	"github.com/MetaDandy/maquetador-angular-backend/config"
	"github.com/MetaDandy/maquetador-angular-backend/src/handlers"
	"github.com/MetaDandy/maquetador-angular-backend/src/repository"
	"github.com/MetaDandy/maquetador-angular-backend/src/services"
)

type Container struct {
	// Users
	UserRepo    *repository.UserRepository
	UserService *services.UserService
	UserHandler *handlers.UserHandler

	// Project
	ProjectRepo    *repository.ProjectRepositories
	ProjectService *services.ProjectService
	ProjectHandler *handlers.ProjectHandler
}

func SetupContainer() *Container {
	// User
	userRepo := repository.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

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
