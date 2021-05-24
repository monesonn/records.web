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

func (q *RecordsQueries) GetGenre(id int) (models.Genre, error) {
	// Define genre variable.
	genre := models.Genre{}

	// Define query string.
	query := `SELECT * FROM genre WHERE genre_id = $1`

	// Send query to database.
	err := q.Get(&genre, query, id)
	if err != nil {
		// Return empty object and error.
		return genre, err
	}

	// Return query result.
	return genre, nil
}

func (q *RecordsQueries) CreateGenre(g *models.Genre) error {
	// Define query string.
	query := `INSERT INTO genre VALUES ($1, $2)`

	// Send query to database.
	_, err := q.Exec(query, g.GenreID, g.Name)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

func (q *RecordsQueries) GetArtists() ([]models.Artist, error) {
	// Define artists variable.
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

func (q *RecordsQueries) GetArtist(id int) (models.Artist, error) {
	// Define artist variable.
	artist := models.Artist{}

	// Define query string.
	query := `SELECT * FROM artist WHERE artist_id = $1`

	// Send query to database.
	err := q.Get(&artist, query, id)
	if err != nil {
		// Return empty object and error.
		return artist, err
	}

	// Return query result.
	return artist, nil
}

func (q *RecordsQueries) CreateArtist(a *models.Artist) error {
	// Define query string.
	query := `INSERT INTO artist VALUES ($1, $2)`

	// Send query to database.
	_, err := q.Exec(query, a.ArtistID, a.Name)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

func (q *RecordsQueries) GetRecords() ([]models.Record, error) {
	// Define artists variable.
	records := []models.Record{}

	query := `SELECT * FROM record`

	// Send query to database.
	err := q.Select(&records, query)
	if err != nil {
		// Return empty object and error.
		return records, err
	}

	// Return query result.
	return records, nil
}

func (q *RecordsQueries) GetRecord(id int) (models.Record, error) {
	// Define artist variable.
	record := models.Record{}

	// Define query string.
	query := `SELECT * FROM record WHERE record_id = $1`

	// Send query to database.
	err := q.Get(&record, query, id)

	if err != nil {
		// Return empty object and error.
		return record, err
	}

	// Return query result.
	return record, nil
}

func (q *RecordsQueries) GetRecordArgs(arg ...interface{}) ([]models.Record, error) {
	// Define artist variable.
	records := []models.Record{}

	// Define query string.
	query := `SELECT * FROM record WHERE `

	var err error
	var artist, genre int
	switch len(arg) {
	case 1:
		artist = arg[0].(int)
		query += `artist_id=$1 `
		err = q.Select(&records, query, artist)
		break
	case 2:
		artist = arg[0].(int)
		genre = arg[1].(int)
		println(artist, genre)
		query += `artist_id=$1 and genre_id=$2 `
		err = q.Select(&records, query, artist, genre)
		break
	default:
		break
	}

	// println(query)
	// Send query to database.

	if err != nil {
		// Return empty object and error.
		return records, err
	}

	// Return query result.
	return records, nil
}

func (q *RecordsQueries) CreateRecord(r *models.Record) error {
	// Define query string.
	query := `INSERT INTO record VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Send query to database.
	_, err := q.Exec(query, r.RecordID, r.ArtistID, r.Title, r.Year, r.Country, r.GenreID, r.LabelID, r.Description)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
