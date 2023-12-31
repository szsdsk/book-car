package controller

import (
	"acs/src/database"
	"acs/src/models"
	"github.com/gofiber/fiber/v2"
)

// GetLocations 获得所有地点数据
func GetLocations(c *fiber.Ctx) error {
	var locations []models.Location
	database.DB.Find(&locations)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    locations,
		"message": "success",
	})
}

// GetLocations 创建地点数据
func CreateLocation(c *fiber.Ctx) error {
	var location models.Location
	if err := c.BodyParser(&location); err != nil {
		return err
	}
	database.DB.Create(&location)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "create successfully",
		"status":  fiber.StatusOK,
	})

}
