// Package controllers is part of MVC pattern.
package controllers

import (
	"fmt"
	"strconv"

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

func GetGenre(c *fiber.Ctx) error {
	// Catch genre ID from URL.
	id, _ := strconv.Atoi(c.Params("id"))

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get genre by ID.
	genre, err := db.GetGenre(id)
	if err != nil {
		// Return, if book not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   fmt.Sprintf("genre with the given ID(%v) is not found.", id),
			"genre": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"genre": genre,
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

func GetArtist(c *fiber.Ctx) error {
	// Catch artist ID from URL.
	id, _ := strconv.Atoi(c.Params("id"))

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get artist by ID.
	artist, err := db.GetArtist(id)
	if err != nil {
		// Return, if book not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    fmt.Sprintf("artist with the given ID(%v) is not found.", id),
			"artist": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"artist": artist,
	})
}

func GetRecords(c *fiber.Ctx) error {
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
	records, err := db.GetRecords()
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
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(records),
		"records": records,
	})
}

func GetRecord(c *fiber.Ctx) error {
	// Catch record ID from URL.
	id, _ := strconv.Atoi(c.Params("id"))

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get record by ID.
	record, err := db.GetRecord(id)
	if err != nil {
		// Return, if book not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     fmt.Sprintf("record with the given ID(%v) is not found.", id),
			"records": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"record": record,
	})
}
