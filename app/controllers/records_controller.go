// Package controllers is part of MVC pattern.
package controllers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/monesonn/records.web/app/models"
	"github.com/monesonn/records.web/pkg/utils"
	"github.com/monesonn/records.web/platform/database"
)

// GetGenres func gets all exists genres.
// @Summary get all exists genres
// @Tags Genre
// @Accept json
// @Produce json
// @Success 200 {array} models.Genre
// @Router /genres [get]
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

// GetGenre func gets genre by given ID or 404 error
// @Summary get genre by given ID
// @Tags Genre
// @Accept json
// @Produce json
// @Param id path string true "Genre ID"
// @Success 200 {object} models.Genre
// @Router /genre/{id} [get]
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
		// Return, if genre not found.
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

// CreateGenre func for creates a new genre.
// @Summary create a new genre
// @Tags Genre
// @Accept json
// @Produce json
// @Success 200 {object} models.Genre
// @Security ApiKeyAuth
// @Router /genre [post]
func CreateGenre(c *fiber.Ctx) error {
	// Create new Genre struct
	genre := &models.Genre{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(genre); err != nil {
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

	// Create a new validator for a Genre model.
	validate := utils.NewValidator()

	// Validate genre fields.
	if err := validate.Struct(genre); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Delete genre by given ID.
	if err := db.CreateGenre(genre); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
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
		// Return, if genre not found.
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

// CreateArtist func for creates a new artist.
// @Summary create a new artist
// @Tags Artist
// @Accept json
// @Produce json
// @Success 200 {object} models.Artist
// @Security ApiKeyAuth
// @Router /artist [post]
func CreateArtist(c *fiber.Ctx) error {
	// Create new Artist struct
	artist := &models.Artist{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(artist); err != nil {
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

	// Create a new validator for a Genre model.
	validate := utils.NewValidator()

	// Validate genre fields.
	if err := validate.Struct(artist); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Delete genre by given ID.
	if err := db.CreateArtist(artist); err != nil {
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
		// Return, if record not found.
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

func GetRecordURL(c *fiber.Ctx) error {
	// Catch records params from URL.
	param := map[string]string{
		"record_id": c.Query("id"),
		"artist_id": c.Query("artist"),
		"genre_id":  c.Query("genre"),
		"year_":     c.Query("year"),
		"title":     c.Query("title"),
		"label":     c.Query("label"),
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

	// Get record by ID.
	record := []models.Record{}

	record, err = db.GetRecordURL(param)

	if err != nil {
		// Return, if record not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     fmt.Sprintf("Record with the given params is not found."),
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

// CreateRecord func for creates a new record.
// @Summary create a new record
// @Tags Record
// @Accept json
// @Produce json
// @Success 200 {object} models.Record
// @Security ApiKeyAuth
// @Router /record [post]
func CreateRecord(c *fiber.Ctx) error {
	// Create new Record struct
	record := &models.Record{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(record); err != nil {
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

	// Create a new validator for a Record model.
	validate := utils.NewValidator()

	// Set initialized default data for record:
	record.Country = sql.NullString{}
	record.Description = sql.NullString{}
	record.LabelID = sql.NullInt32{}
	record.Year = sql.NullInt32{}

	// Validate genre fields.
	if err := validate.Struct(record); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Delete record by given ID.
	if err := db.CreateRecord(record); err != nil {
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
		"record": record,
	})
}
