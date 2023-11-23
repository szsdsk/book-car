package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func MyLog(c *fiber.Ctx) error {
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
}
