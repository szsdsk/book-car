package controller

import (
	"acs/src/database"
	"acs/src/models"
	"github.com/gofiber/fiber/v2"
)

// RenderCars 返回所有车辆数据，将数据返回给前端，用于渲染模板引擎。
func RenderCars(c *fiber.Ctx) error {
	var cars []models.Car
	database.DB.Order("capacity").Find(&cars)
	return c.Status(fiber.StatusOK).Render("index", fiber.Map{
		"Cars": cars,
	})
}

// FilterCars 返回座位数大于等于所给值车辆数据
func FilterCars(c *fiber.Ctx) error {
	num := c.Params("num")
	var cars []models.Car
	if num == "" {
		database.DB.Find(&cars)
	} else {
		database.DB.Where("capacity >= ?", num).Order("capacity").Find(&cars)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Cars": cars,
	})
}

// Admin admin界面的三个表，返回给前端渲染。
func Admin(c *fiber.Ctx) error {
	type Probation struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		EmailAddress string `json:"email"`
	}
	var probation []Probation
	database.DB.Raw(`SELECT * FROM customers_on_probation()`).Scan(&probation)

	type Popular struct {
		LocationId      int    `json:"locationId"`
		StreetAddress   string `json:"stressAddress"`
		TelePhone       string `json:"telePhone"`
		NumberOfRentals int    `json:"numberOfRentals"`
	}
	var results []Popular
	database.DB.Raw(`SELECT * FROM popular_locations()`).Scan(&results)

	type Trends struct {
		Make_               string `json:"make"`
		Model_              string `json:"model"`
		IsStudent_          bool   `json:"isStudent"`
		NumberOfTimesRented int    `json:"numberOfTimesRented"`
	}
	var trends []Trends
	database.DB.Raw(`SELECT * FROM rental_trends()`).Scan(&trends)

	return c.Status(fiber.StatusOK).Render("admin", fiber.Map{
		"customers": probation,
		"locations": results,
		"trends":    trends,
	})
}

func BookingRender(c *fiber.Ctx) error {
	carId := c.Params("id")
	var locations []models.Location
	database.DB.Find(&locations)
	return c.Status(fiber.StatusOK).Render("booking", fiber.Map{
		"carId":     carId,
		"locations": locations,
	})
}
