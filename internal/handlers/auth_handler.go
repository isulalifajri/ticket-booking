package handlers

import (
	"ticket-booking/internal/dto"
	"ticket-booking/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(
	authService *services.AuthService,
) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// register handler
func (h *AuthHandler) Register(c *fiber.Ctx) error {

	var req dto.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	err := h.authService.Register(req)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "register success",
	})
}

// login handler
func (h *AuthHandler) Login(c *fiber.Ctx) error {

	var req dto.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	token, err := h.authService.Login(req)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "login success",
		"token":   token,
	})

	// return c.JSON(fiber.Map{
	// 	"message": "login success",
	// 	"user": fiber.Map{
	// 		"id":    user.ID,
	// 		"name":  user.Name,
	// 		"email": user.Email,
	// 		"role":  user.Role,
	// 	},
	// })
}