package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monesonn/records.web/platform/database"
)

func RenderProducts(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all records
	product, err := db.GetProductView()
	if err != nil {
		// Return, if records not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "records were not found",
			"count":   0,
			"records": nil,
			"err":     err.Error(),
		})
	}
	return c.Render("explore", fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(product),
		"product": product,
	})
}

func RenderAdminPanel(c *fiber.Ctx) error {
	// Create database connection.

	// db, err := database.OpenDBConnection()
	// if err != nil {
	// 	// Return status 500 and database connection error.
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	return c.Render("admin", fiber.Map{})
}
