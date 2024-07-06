package controller

import "github.com/gofiber/fiber/v2"

func BlogList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Blog List",
	}

	c.Status(200)
	return c.JSON(context)
}

// return status code 201
func BlogCreate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Add Blog",
	}

	c.Status(201)
	return c.JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Update Blog",
	}

	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Delete Blog for the given ID",
	}

	c.Status(200)
	return c.JSON(context)
}