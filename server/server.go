package main

import (
	"github.com/JeaMianh/MyBlog/database"
	"github.com/JeaMianh/MyBlog/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// 在 main 之前自动调用 init 函数，无需显式调用
func init() {

	database.ConnectDB()
}

func main() {
	// DB() 是 *gorm.DB 类型的一个方法，用于获取底层的 *sql.DB 对象 sqlDb
	// 这里的底层是相对于 *gorm.DB 类型的高级别(ORM)而言的
	// 高层操作就是对数据对象的 CRUD
	// 底层就是直接对数据库的操作，例如创建、复用、断开，也可以直接用 SQL 语句
	sqlDb, err := database.DBconn.DB()

	if err != nil {
		panic("Error in sql connection.")
	}

	// main 结束后执行
	defer sqlDb.Close()

	app := fiber.New()

	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":5555")
}
