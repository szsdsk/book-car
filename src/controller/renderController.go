package controller

import (
	"acs/src/database"
	"acs/src/models"
	"github.com/gofiber/fiber/v2"
)

func RenderCars(c *fiber.Ctx) error {
	var cars []models.Car
	database.DB.Order("capacity").Find(&cars)
	return c.Status(fiber.StatusOK).Render("index", fiber.Map{
		"Cars": cars,
	})
}

func FilterCars(c *fiber.Ctx) error {
	num := c.Params("num")
	var cars []models.Car
	if num == "" {
		database.DB.Find(&cars)
	} else {
		database.DB.Where("capacity >= ?", num).Order("capacity").Find(&cars)
	}

	//return c.Status(http.StatusOK).Render("index", fiber.Map{
	//	"Cars": cars,
	//})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Cars": cars,
	})
}
