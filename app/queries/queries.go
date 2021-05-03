// Package queries is used for represent view of data.
package queries

import (
	"github.com/jmoiron/sqlx"
	"github.com/monesonn/records.web/app/models"
)

type RecordsQueries struct {
	*sqlx.DB
}

func (q *RecordsQueries) GetGenres() ([]models.Genre, error) {
	// Define review variable.
	genres := []models.Genre{}

	query := `SELECT * FROM genre`

	// Send query to database.
	err := q.Select(&genres, query)
	if err != nil {
		// Return empty object and error.
		return genres, err
	}

	// Return query result.
	return genres, nil
}

func (q *RecordsQueries) GetArtists() ([]models.Artist, error) {
	// Define review variable.
	artists := []models.Artist{}

	query := `SELECT * FROM artist`

	// Send query to database.
	err := q.Select(&artists, query)
	if err != nil {
		// Return empty object and error.
		return artists, err
	}

	// Return query result.
	return artists, nil
}
