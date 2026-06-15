package routes

import (
	"ticket-booking/internal/handlers"
	"ticket-booking/internal/middleware"
	"ticket-booking/internal/repositories"
	"ticket-booking/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(
	app *fiber.App,
	db *gorm.DB,
) {

	// Repository
	userRepo := repositories.NewUserRepository(db)
	movieRepo := repositories.NewMovieRepository(db)

	// Service
	authService := services.NewAuthService(userRepo)
	movieService := services.NewMovieService(movieRepo)

	// Handler
	authHandler := handlers.NewAuthHandler(authService)
	movieHandler := handlers.NewMovieHandler(movieService)

	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	api := app.Group("/api")

	// Public Routes
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	// Protected Routes
	protected := api.Group(
		"",
		middleware.AuthMiddleware(),
	)

	protected.Get("/profile", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"user_id": c.Locals("user_id"),
			"email":   c.Locals("email"),
			"role":    c.Locals("role"),
		})
	})

	// Movie Routes
	protected.Post("/movies", movieHandler.Create)
	protected.Get("/movies", movieHandler.GetAll)
	protected.Get("/movies/:id", movieHandler.GetByID)
	protected.Put("/movies/:id", movieHandler.Update)
	protected.Delete("/movies/:id", movieHandler.Delete)
}