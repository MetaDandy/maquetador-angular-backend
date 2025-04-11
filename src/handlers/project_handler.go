package handlers

import (
	"github.com/MetaDandy/maquetador-angular-backend/src/dtos"
	"github.com/MetaDandy/maquetador-angular-backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ProjectHandler struct {
	projectService *services.ProjectService
}

func NewProjectHandler(projectService *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

func (h *ProjectHandler) RegisterProjectRoutes(router fiber.Router) {
	router.Post("/project/register", h.CreateProject)
}

func (h *ProjectHandler) CreateProject(c *fiber.Ctx) error {
	var req dtos.ProjectCreate
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	project, err := h.projectService.CreateProject(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create project",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    project,
	})
}
