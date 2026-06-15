package handlers

import (
	"strconv"
	"errors"
	"gorm.io/gorm"

	"ticket-booking/internal/dto"
	"ticket-booking/internal/services"

	"github.com/gofiber/fiber/v2"
)

type MovieHandler struct {
	movieService *services.MovieService
}

func NewMovieHandler(
	movieService *services.MovieService,
) *MovieHandler {
	return &MovieHandler{
		movieService: movieService,
	}
}

func (h *MovieHandler) Create(c *fiber.Ctx) error {

	var req dto.CreateMovieRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	err := h.movieService.Create(req)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "movie created successfully",
	})
}

func (h *MovieHandler) GetAll(c *fiber.Ctx) error {

	movies, err := h.movieService.GetAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": movies,
	})
}

func (h *MovieHandler) GetByID(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(
		c.Params("id"),
		10,
		64,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	movie, err := h.movieService.GetByID(uint(id))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "movie not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": movie,
	})
}

func (h *MovieHandler) Update(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(
		c.Params("id"),
		10,
		64,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	var req dto.UpdateMovieRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	err = h.movieService.Update(uint(id), req)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "movie updated successfully",
	})
}

func (h *MovieHandler) Delete(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(
		c.Params("id"),
		10,
		64,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	err = h.movieService.Delete(uint(id))

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "movie not found",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "movie deleted successfully",
	})
}