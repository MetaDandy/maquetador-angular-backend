package handlers

import (
	"github.com/MetaDandy/maquetador-angular-backend/pkg"
	"github.com/MetaDandy/maquetador-angular-backend/src/dtos"
	"github.com/MetaDandy/maquetador-angular-backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterUserRoutes(router fiber.Router) {
	router.Post("/user/register", h.CreateUser)
	router.Post("/login", h.Login)
	router.Get("/", h.FindAll)
	router.Get("/:id", h.FindById)
	router.Delete("/:id", h.Delete)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dtos.UserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
			"err":   err,
		})
	}

	user, err := h.userService.CreateUser(req)
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

func (h *UserHandler) FindAll(c *fiber.Ctx) error {
	opts := pkg.NewFindAllOptionsFromQuery(c)

	users, err := h.userService.GetAllUsers(opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to obtein users",
			"err":   err,
		})
	}

	return c.JSON(users)
}

func (h *UserHandler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userService.FindUserById(id)
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

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userService.Delete(id)
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

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dtos.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	loginResponse, err := h.userService.Login(req)
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
