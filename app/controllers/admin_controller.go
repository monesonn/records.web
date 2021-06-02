package controllers

import (
	"github.com/gofiber/fiber/v2"
)

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

	return c.Render("admin", "text/html")
}
