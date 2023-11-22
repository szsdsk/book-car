package controller

import (
	"acs/src/database"
	"acs/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"net/http"
	"time"
)

func GetCars(c *fiber.Ctx) error {
	carId := c.Query("carId")
	maker := c.Query("make")
	model := c.Query("model")
	pricePerHour := c.Query("pricePerHour")
	pricePerDay := c.Query("pricePerDay")
	capacity := c.Query("capacity")
	query := database.DB
	if carId != "" {
		query = query.Where("car_id = ?", carId)
	}
	if maker != "" {
		query = query.Where("make = ?", maker)
	}
	if model != "" {
		query = query.Where("model = ?", model)
	}
	if pricePerHour != "" {
		query = query.Where("price_per_hour = ?", pricePerHour)
	}
	if pricePerDay != "" {
		query = query.Where("price_per_day = ?", pricePerDay)
	}
	if capacity != "" {
		query = query.Where("capacity = ?", capacity)
	}
	var cars []models.Car
	query.Find(&cars)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success",
		"status":  http.StatusOK,
		"data":    cars,
	})
}

func CreateCar(c *fiber.Ctx) error {
	var car models.Car
	// 传入car的指针，将传来的JSON赋值给car对象。
	if err := c.BodyParser(&car); err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).SendString("Invalid request")
	}
	//赋值创建和更新时间。
	car.CreateAt = time.Now()
	car.UpdateAt = time.Now()
	//数据库根据car对象在表中添加，相当于insert values into table
	result := database.DB.Create(&car)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error,
			"status":  http.StatusBadRequest,
			"data":    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create successfully",
		"status":  http.StatusOK,
		"data":    car,
	})
}

func UpdateCar(c *fiber.Ctx) error {
	id := c.Params("id")
	var car models.Car
	database.DB.Where("car_id = ?", id).First(&car)
	if err := c.BodyParser(&car); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request")
	}
	car.UpdateAt = time.Now()
	database.DB.Model(&car).Updates(car)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "modify successfully",
		"status":  http.StatusOK,
		"data":    car,
	})
}

func DeleteCar(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.Car{}, id)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
		"status":  http.StatusOK,
	})
}
