package controller

import (
	"log"

	"github.com/JeaMianh/MyBlog/database"
	"github.com/JeaMianh/MyBlog/model"
	"github.com/gofiber/fiber/v2"
)

func BlogList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"message":    "Blog List",
	}

	var records []model.Blog

	// 查询到的结果直接赋值给 records
	result := database.DBconn.Find(&records)
	if result.Error != nil {

		// 打印后 exit(1)
		log.Println("Error in fetching records.", result.Error)
		context["statusText"] = "Error"
		// 服务器请求错误
		return c.Status(500).JSON(context)
	}

	context["blog_records"] = records

	return c.Status(200).JSON(context)
}

// return status code 201
func BlogCreate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Error",
		"message":    "Add Blog",
	}

	// 返回一个指针
	record := new(model.Blog)

	// 解析请求体中的数据，映射到 record 的指针中。之所以是指针，是因为需要进行修改。
	err := c.BodyParser(record)

	if err != nil {
		log.Println("Error in parsing request.", err)

		// 400 Bad Request 请求格式不正确
		return c.Status(400).JSON(context)
	}

	result := database.DBconn.Create(record)
	if result.Error != nil {
		log.Println("Error in saving data.", result.Error)

		// 500 服务器内部错误
		return c.Status(500).JSON(context)
	}

	context["statusText"] = "OK"
	context["message"] = "Record is saved successfully."
	context["data"] = record

	return c.Status(201).JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Error",
		"message":    "Update Blog",
	}

	// 获取 ID 参数
	id := c.Params("id")

	// 解析更新数据
	updatedRecord := new(model.Blog)
	err := c.BodyParser(updatedRecord)
	if err != nil {
		log.Println("Error in parsing request.", err)

		// 400 Bad Request 请求格式不正确
		return c.Status(400).JSON(context)
	}

	// 检查 ID 是否合法
	var existingRecord model.Blog
	result := database.DBconn.First(&existingRecord, id)
	if result.Error != nil {
		log.Println("Error in finding record.", result.Error)

		// Record not found
		return c.Status(404).JSON(context)
	}

	// 更新数据
	existingRecord.Title = updatedRecord.Title
	existingRecord.Content = updatedRecord.Content

	// 保存数据
	result = database.DBconn.Save(&existingRecord)
	if result.Error != nil {
		log.Println("Error in saving data.", result.Error)

		// 500 服务器内部错误
		return c.Status(500).JSON(context)
	}

	// 返回成功响应
	context["statusText"] = "OK"
	context["message"] = "Record updated successfully"
	context["data"] = existingRecord

	return c.Status(200).JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Error",
		"message":    "Delete Blog for the given ID",
	}

	id := c.Params("id")

	var record model.Blog
	result := database.DBconn.First(&record, id)
	if result.Error != nil {
		log.Println("Error in finding record.", result.Error)

		// Record not found
		return c.Status(404).JSON(context)
	}

	if result := database.DBconn.Delete(&record); result.Error != nil {
		log.Println("Error in deleting record.", result.Error)
		context["message"] = "Delete error"
		return c.Status(500).JSON(context)
	}

	context["statusText"] = "OK"
	context["message"] = "Record deleted successfully"

	return c.Status(200).JSON(context)
}
