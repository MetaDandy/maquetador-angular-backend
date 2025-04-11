package handlers

import (
	"github.com/MetaDandy/maquetador-angular-backend/pkg"
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
	router.Get("/", h.FindAll)
	router.Get("/:id", h.FindById)
	router.Delete("/:id", h.Delete)
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

func (h *ProjectHandler) FindAll(c *fiber.Ctx) error {
	opts := pkg.NewFindAllOptionsFromQuery(c)

	projects, err := h.projectService.GetAll(opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to obtein projects",
			"err":   err,
		})
	}

	return c.JSON(projects)
}

func (h *ProjectHandler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.projectService.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to find project",
			"err":   err,
		})
	}
	if user == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No project finded",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Project finded successfully",
		"data":    user,
	})
}

func (h *ProjectHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	project, err := h.projectService.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to find project",
			"err":   err,
		})
	}
	if project == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No project finded",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Project deleted successfully",
		"data":    project,
	})
}
