package routes

import (
	"ticket-booking/internal/handlers"
	"ticket-booking/internal/repositories"
	"ticket-booking/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	
	"ticket-booking/internal/middleware"
)

func SetupRoutes(
	app *fiber.App,
	db *gorm.DB,
) {

	userRepo := repositories.NewUserRepository(db)

	authService := services.NewAuthService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	api := app.Group("/api")

	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	api.Get(
		"/profile",
		middleware.AuthMiddleware(),
		func(c *fiber.Ctx) error {

			return c.JSON(fiber.Map{
				"user_id": c.Locals("user_id"),
				"email":   c.Locals("email"),
				"role":    c.Locals("role"),
			})
		},
	)
}