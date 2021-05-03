// Package controllers is part of MVC pattern.
package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monesonn/records.web/platform/database"
)

func GetGenres(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all genres.
	genres, err := db.GetGenres()
	if err != nil {
		// Return, if genres not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "genres were not found",
			"count":  0,
			"genres": nil,
		})
	}
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"count":  len(genres),
		"genres": genres,
	})
}

func GetArtists(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all artists.
	artists, err := db.GetArtists()
	if err != nil {
		// Return, if artists not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "artists were not found",
			"count":   0,
			"artists": nil,
		})
	}
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(artists),
		"artists": artists,
	})
}
