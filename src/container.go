package src

import (
	"github.com/MetaDandy/maquetador-angular-backend/config"
	"github.com/MetaDandy/maquetador-angular-backend/src/modules/project"
	"github.com/MetaDandy/maquetador-angular-backend/src/modules/user"
)

type Container struct {
	// Users
	UserRepo    *user.Repository
	UserService *user.Service
	UserHandler *user.Handler

	// Project
	ProjectRepo    *project.Repository
	ProjectService *project.Service
	ProjectHandler *project.Handler
}

func SetupContainer() *Container {
	// User
	userRepo := user.NewRepository(config.DB)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	// Project
	projectRepo := project.NewRepository(config.DB)
	projectService := project.NewService(projectRepo, userRepo)
	projectHandler := project.NewHandler(projectService)

	return &Container{
		// User
		UserRepo:    userRepo,
		UserService: userService,
		UserHandler: userHandler,

		// Project
		ProjectRepo:    projectRepo,
		ProjectService: projectService,
		ProjectHandler: projectHandler,
	}
}
