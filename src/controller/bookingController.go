package controller

import (
	"acs/src/database"
	"acs/src/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"math"
	"time"
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

func SubmitBookRecord(c *fiber.Ctx) error {
	var record models.BookRecord
	var user models.Customer
	var car models.Car
	if err := c.BodyParser(&record); err != nil {
		log.Fatal(err)
		return err
	}
	if err := c.BodyParser(&user); err != nil {
		log.Fatal(err)
		return err
	}
	database.DB.Model(&models.Car{}).Where("id = ?", record.CarId).First(&car)
	record.PricePerHour = car.PricePerHour
	record.PricePerDay = car.PricePerDay
	record.CustomerId = user.Uid
	var count int64
	database.DB.Model(&models.BookRecord{}).Count(&count)
	record.ReservationNum = time.Now().Format("Jan") + fmt.Sprintf("-%03d", count)
	if database.DB.Where("uid = ?", user.Uid).First(&models.Customer{}).RowsAffected == 0 {
		database.DB.Create(&user)
	}
	record.ReservedDate = time.Now()
	database.DB.Create(&record)
	Hours := math.Ceil(record.DropOfTime.Sub(record.PickUpTime).Hours())
	price := math.Floor(Hours/24)*record.PricePerDay + float64(int(Hours)%24)*record.PricePerHour
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "book successfully",
		"data": fiber.Map{
			"price": price,
		},
		"status": fiber.StatusOK,
	})
}
