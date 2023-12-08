package controller

import (
	"acs/src/database"
	"acs/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetBookRecords(c *fiber.Ctx) error {
	var bookRecodes []models.BookRecord
	database.DB.Preload("Car").Preload("Customer").Preload("Location").Find(&bookRecodes)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    bookRecodes,
		"message": "success",
	})
}

func CreateBookRecord(c *fiber.Ctx) error {
	var record models.BookRecord
	if err := c.BodyParser(&record); err != nil {
		return err
	}
	database.DB.Create(&record)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "create successfully",
		"status":  fiber.StatusOK,
	})
}

//func UpdateBookRecord(c *fiber.Ctx) {
//	var record models.BookRecord
//
//}
