package user

import (
	"github.com/MetaDandy/maquetador-angular-backend/helper"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterUserRoutes(router fiber.Router) {
	grp := router.Group("/users")
	grp.Post("/register", h.CreateUser)
	grp.Post("/login", h.Login)
	grp.Get("/", h.FindAll)
	grp.Get("/:id", h.FindById)
	grp.Delete("/:id", h.Delete)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var req UserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
			"err":   err,
		})
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
			"err":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    user,
	})
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	users, err := h.service.GetAllUsers(opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to obtein users",
			"err":   err,
		})
	}

	return c.JSON(users)
}

func (h *Handler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.FindUserById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to find user",
			"err":   err,
		})
	}
	if user == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No user finded",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User finded successfully",
		"data":    user,
	})
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to find user",
			"err":   err,
		})
	}
	if user == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No user finded",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User deleted successfully",
		"data":    user,
	})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	loginResponse, err := h.service.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   loginResponse.Token,
	})
}
