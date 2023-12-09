package routes

import (
	"acs/src/controller"
	"acs/src/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"net/http"
)

func createMyRender() *html.Engine {
	engine := html.New("./templates", ".html")
	return engine
}

func InitRouter() {
	//创建app实例
	app := fiber.New(fiber.Config{
		Views: createMyRender(), //渲染html模板
	})
	// 使用默认的请求日志中间件
	app.Use(middleware.MyLog)

	//引入静态文件夹
	app.Static("/static", "./static")
	app.Get("/", func(c *fiber.Ctx) error {
		//返回html模板渲染的index.html
		return c.Redirect("/index", http.StatusPermanentRedirect)
	})

	app.Get("/index", controller.RenderCars)

	app.Get("/index/:num", controller.FilterCars)

	api := app.Group("/api")
	{
		car := api.Group("/cars")
		{
			car.Get("/", controller.GetCars)
			car.Post("/", controller.CreateCar)
			car.Put("/:id", controller.UpdateCar)
			car.Delete("/:id", controller.DeleteCar)
		}

		booking := api.Group("/bookings")
		{
			booking.Get("/", controller.GetBookRecords)
			booking.Post("/", controller.CreateBookRecord)
		}

		location := api.Group("/locations")
		{
			location.Get("/", controller.GetLocations)
			location.Post("/", controller.CreateLocation)
		}

		customer := api.Group("/customers")
		{
			customer.Get("/", controller.GetCustomers)
			customer.Post("/", controller.GetCustomers)
		}
	}
	admin := app.Group("/admin")
	{
		admin.Get("/", controller.Admin)
		admin.Get("/popular", controller.PopularLocatoins)
		admin.Get("/trends", controller.RentalTrends)
		admin.Patch("/increase", controller.RentalIncrease)
	}
	booking := app.Group("booking")
	{
		booking.Get("/:id", controller.BookingRender)
		booking.Post("/submit", controller.SubmitBookRecord)
	}
	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
