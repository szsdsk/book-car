package controller

import (
	"acs/src/database"
	"acs/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetCustomers(c *fiber.Ctx) error {
	var customers []models.Customer
	database.DB.Find(&customers)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    customers,
		"message": "success",
	})
}

func CreateCustomer(c *fiber.Ctx) error {
	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		return err
	}
	database.DB.Create(&customer)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "create successfully",
		"status":  fiber.StatusOK,
	})

}
