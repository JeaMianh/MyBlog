package router

import (
	"github.com/JeaMianh/MyBlog/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// 路径，访问处理函数，返回一个 JSON 响应
	app.Get("/", controller.BlogList)
	app.Post("/", controller.BlogCreate)
	app.Put("/:id", controller.BlogUpdate)
	app.Delete("/:id", controller.BlogDelete)
}
