package routes

import (
	"acs/src/controller"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
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
	app.Use(func(c *fiber.Ctx) error {
		// 获取方法
		method := c.Method()

		// 根据方法设置颜色
		var color string
		switch method {
		case fiber.MethodGet:
			color = "\x1b[32m" // 绿色
		case fiber.MethodPost:
			color = "\x1b[34m" // 蓝色
		case fiber.MethodPut:
			color = "\x1b[33m" // 黄色
		case fiber.MethodDelete:
			color = "\x1b[31m" // 红色
		default:
			color = "\x1b[0m" // 默认颜色
		}

		// 输出日志
		log.Println(fmt.Sprintf("%s[%s] %s %s\x1b[0m", color, method, c.Path(), c.IP()))

		// 继续处理请求
		return c.Next()
	})

	//引入静态文件夹
	app.Static("/static", "./static")
	app.Get("/", func(c *fiber.Ctx) error {
		//返回html模板渲染的index.html
		return c.Redirect("/index", http.StatusPermanentRedirect)
	})

	app.Get("/index", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/:num", controller.NumberOfPassengers)

	api := app.Group("/api")
	{
		api.Get("/cars", controller.GetCars)
		api.Post("/cars", controller.CreateCar)
		api.Put("/cars/:id", controller.UpdateCar)
		api.Delete("/cars/:id", controller.DeleteCar)
	}
	admin := app.Group("/admin")
	{
		admin.Get("/popular", controller.PopularLocatoins)
		admin.Get("/trends", controller.RentalTrends)
		admin.Post("/increase", controller.RentalIncrease)
	}
	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
