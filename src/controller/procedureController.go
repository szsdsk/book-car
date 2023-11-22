package controller

import (
	"acs/src/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func CustomersOnProbation(c *fiber.Ctx) error {
	type Probation struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		EmailAddress string `json:"email"`
	}
	var probation []Probation
	database.DB.Raw(`SELECT * FROM customers_on_probation()`).Scan(&probation)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"data":    probation,
		"message": "OK",
	})
}

func NumberOfPassengers(c *fiber.Ctx) error {
	num := c.Params("num")
	res, _ := strconv.Atoi(num)
	type Car struct {
		CarId              int     `json:"id"`
		CarMake            string  `json:"make"`
		CarModel           string  `json:"model"`
		PricePerhour       float64 `json:"pricePerHour"`
		NumberOfPassengers int     `json:"numberOfPassengers"`
	}
	var cars []Car
	query := fmt.Sprintf("SELECT * FROM number_of_passengers(%d)", res)
	database.DB.Raw(query).Scan(&cars)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"data":    cars,
		"message": "OK",
	})
}

func PopularLocatoins(c *fiber.Ctx) error {
	//TABLE(location_id bigint, street_address character
	//varying, tele_phone character varying, number_of_rentals bigint)
	type Popular struct {
		LocationId      int    `json:"locationId"`
		StreetAddress   string `json:"stressAddress"`
		TelePhone       string `json:"telePhone"`
		NumberOfRentals int    `json:"numberOfRentals"`
	}
	var results []Popular
	database.DB.Raw(`SELECT * FROM popular_locations()`).Scan(&results)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"data":    results,
		"message": "OK",
	})
}

func RentalTrends(c *fiber.Ctx) error {
	//TABLE(make_ character varying, model_ character varying,
	//"isStudent?" boolean, number_of_times_rented bigint)
	type Trends struct {
		Make_               string `json:"make"`
		Model_              string `json:"model"`
		IsStudent_          bool   `json:"isStudent"`
		NumberOfTimesRented int    `json:"numberOfTimesRented"`
	}
	var results []Trends
	database.DB.Raw(`SELECT * FROM rental_trends()`).Scan(&results)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"data":    results,
		"message": "OK",
	})
}

func RentalIncrease(c *fiber.Ctx) error {
	database.DB.Exec(`CALL rental_increase()`)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "OK",
	})
}
