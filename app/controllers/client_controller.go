package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/monesonn/records.web/app/models"
	"github.com/monesonn/records.web/pkg/utils"
	"github.com/monesonn/records.web/platform/database"
)

func GetClientByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	client, err := db.GetClientByUUID(uuid)
	if err != nil {
		// Return, if client not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "Client were not found.",
			"client": nil,
		})
	}
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"client": client,
	})
}

func GetClientByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	client, err := db.GetClientByEmail(email)
	if err != nil {
		// Return, if client not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "Client were not found.",
			"client": nil,
		})
	}
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"client": client,
	})
}

func GetUser(c *fiber.Ctx) error {
	username := c.Params("username")

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	client, err := db.GetUser(username)
	if err != nil {
		// Return, if client not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "Client were not found.",
			"client": nil,
		})
	}
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"client": client,
	})
}

func CreateProfile(c *fiber.Ctx) error {
	profile := &models.Profile{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(profile); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
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

	validate := utils.NewValidator()

	if err := validate.Struct(profile); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	if err := db.CreateProfile(profile); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"client": profile,
	})
}

func GetReview(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	review, err := db.GetReview(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    fmt.Sprintf("comments with the given product id(%v) is not found.", id),
			"review": nil,
		})
	}

	return c.JSON(fiber.Map{
		"review": review,
	})
}
