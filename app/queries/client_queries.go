package queries

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/monesonn/records.web/app/models"
)

type ClientQueries struct {
	*sqlx.DB
}

func (q *RecordsQueries) GetClientByUUID(uuid string) (models.Client, error) {
	client := models.Client{}
	query := `SELECT * FROM clients_view where id = $1`

	err := q.Get(&client, query, uuid)
	if err != nil {
		fmt.Println(err)
		return client, err
	}
	return client, nil
}

func (q *RecordsQueries) GetClientByEmail(email string) (models.Client, error) {
	client := models.Client{}
	query := `SELECT * FROM clients_view where email = $1`

	fmt.Println(email)

	err := q.Get(&client, query, email)
	if err != nil {
		fmt.Println(err)
		return client, err
	}
	return client, nil
}

func (q *RecordsQueries) GetUser(username string) (models.Client, error) {
	client := models.Client{}
	query := `SELECT * FROM clients_view where username = $1`

	err := q.Get(&client, query, username)
	if err != nil {
		fmt.Println(err)
		return client, err
	}
	return client, nil
}

func (q *RecordsQueries) CreateProfile(c *models.Profile) error {
	// Define query string.
	query := `INSERT INTO profile(user_id, first_name,last_name,gender,birthday) VALUES ($1, $2, $3, $4, $5)`

	// Send query to database.
	_, err := q.Exec(query, c.ID, c.FirstName, c.LastName, c.Gender, c.Birthday)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

func (q *RecordsQueries) GetReview(id int) ([]models.Review, error) {
	reviews := []models.Review{}
	query := `SELECT * FROM review where view_product_id = $1`
	err := q.Select(&reviews, query, id)
	if err != nil {
		return reviews, err
	}
	return reviews, nil
}
