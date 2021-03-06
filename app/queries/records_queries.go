// Package queries is used for represent view of data.
package queries

import (
	"fmt"
	"regexp"

	"github.com/jmoiron/sqlx"
	"github.com/monesonn/records.web/app/models"
)

type RecordsQueries struct {
	*sqlx.DB
}

func (q *RecordsQueries) GetGenres() ([]models.Genre, error) {
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
	// Define record variable.
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

func (q *RecordsQueries) GetRecordURL(param map[string]string) ([]models.Record, error) {
	// Define records variable.
	records := []models.Record{}

	// Define base of SQL query string.
	query := `SELECT * FROM record WHERE `

	// Go through map & add param to SQL query if not empty
	for k, v := range param {
		if len(v) != 0 {
			query += fmt.Sprintf("%v in (%v) and ", k, v)
		}
	}

	// So, as it not regular task
	// It's have kinda bad implementation, this 2 lines
	// Delete last 'and' in SQL query
	re, _ := regexp.Compile(`\sand\s$`)
	query = re.ReplaceAllString(query, "")

	// Send query to database.
	err := q.Select(&records, query)

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

func (q *RecordsQueries) GetProducts() ([]models.ProductView, error) {
	products := []models.ProductView{}

	query := `SELECT * FROM records_view`

	// Send query to database.
	err := q.Select(&products, query)
	if err != nil {
		// Return empty object and error.
		return products, err
	}

	// Return query result.
	return products, nil
}

func (q *RecordsQueries) GetProductURL(param map[string]string) ([]models.Product, error) {
	// Define records variable.
	product := []models.Product{}

	// Define base of SQL query string.
	query := `SELECT * FROM products WHERE `

	// Go through map & add param to SQL query if not empty
	for k, v := range param {
		if len(v) != 0 {
			if k == "medium_type" {
				re, _ := regexp.Compile(`,`)
				v = re.ReplaceAllString(v, "','")
				query += fmt.Sprintf("%v in ('%v') and ", k, v)
			} else {
				query += fmt.Sprintf("%v in (%v) and ", k, v)

			}
		}
	}

	re, _ := regexp.Compile(`\sand\s$`)
	query = re.ReplaceAllString(query, "")

	println(query)
	// Send query to database.
	err := q.Select(&product, query)

	if err != nil {
		// Return empty object and error.
		return product, err
	}

	// Return query result.
	return product, nil
}

func (q *RecordsQueries) GetProductView() ([]models.ProductView, error) {
	// Define record variable.
	product := []models.ProductView{}

	// Define query string.
	query := `SELECT * FROM records_view`

	// Send query to database.
	err := q.Select(&product, query)

	if err != nil {
		// Return empty object and error.
		return product, err
	}

	// Return query result.
	return product, nil
}
