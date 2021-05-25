package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/monesonn/records.web/pkg/repository"
	"github.com/monesonn/records.web/pkg/utils"
	"github.com/monesonn/records.web/platform/database"
)

func GetClients(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Set credential `book:update` from JWT data of current book.
	credential := claims.Credentials[repository.ClientGetCredential]

	// Only book creator with `book:update` credential can update his book.
	if !credential {
		// Return status 403 and permission denied error message.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check credentials of your token",
		})
	}
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all clients.
	clients, err := db.GetClients()
	if err != nil {
		// Return, if genres not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "clients were not found",
			"count":  0,
			"genres": nil,
		})
	}
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"count":  len(clients),
		"genres": clients,
	})
}
