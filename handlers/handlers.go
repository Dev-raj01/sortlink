package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/Dev-raj01/sortlink/database"
	"github.com/Dev-raj01/sortlink/models"
)

func InitRouter(app *fiber.App) {
	short := app.Group("/api/v1")

	short.Get("/shorts", GetAllURLs)
	short.Get("/shorts/:id", GetAURLs)
	short.Post("/shorts", CreateURL)
}

func GetAllURLs(c *fiber.Ctx) error {
	URLs, err := database.GetAllURLs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "cannot get all urls " + err.Error(),
		})
	}
	return c.JSON(URLs)
}

func GetAURLs(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing id " + err.Error(),
		})
	}
	URL, err := database.GetURL(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting url from database " + err.Error(),
		})
	}
	return c.JSON(URL)
}

func CreateURL(c *fiber.Ctx) error {
	var URL models.URL
	if err := c.BodyParser(&URL); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing body " + err.Error(),
		})
	}
	if err := database.CreateURL(URL); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error creating record for URL " + err.Error(),
		})
	}
	return c.JSON(URL)
}
