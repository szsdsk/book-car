package controller

import (
	"acs/src/database"
	"acs/src/models"
	"github.com/gofiber/fiber/v2"
)

func RenderCars(c *fiber.Ctx) error {
	var cars []models.Car
	database.DB.Find(&cars)
	return c.Status(fiber.StatusOK).Render("index", fiber.Map{
		"Cars": cars,
	})
}
