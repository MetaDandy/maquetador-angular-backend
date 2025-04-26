package project

import (
	"github.com/MetaDandy/maquetador-angular-backend/helper"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) RegisterProjectRoutes(router fiber.Router) {
	grp := router.Group("/projects")
	grp.Post("/create", h.CreateProject)
	grp.Get("/", h.FindAll)
	grp.Get("/owner/:id", h.FindAllProjectsByUser)
	grp.Get("/:id", h.FindById)
	grp.Delete("/:id", h.Delete)
}

func (h *Handler) CreateProject(c *fiber.Ctx) error {
	var req ProjectCreate
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	project, err := h.svc.CreateProject(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create project",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Project created successfully",
		"data":    project,
	})
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	projects, err := h.svc.GetAll(opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to obtein projects",
			"err":   err,
		})
	}

	return c.JSON(projects)
}

func (h *Handler) FindAllProjectsByUser(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)
	id := c.Params("id")

	projects, err := h.svc.FindAllProjectsByUser(id, opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to obtein projects",
			"err":   err,
		})
	}

	return c.JSON(projects)
}

func (h *Handler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	project, err := h.svc.FindByID(id)
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
		"message": "Project finded successfully",
		"data":    project,
	})
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	project, err := h.svc.Delete(id)
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
