package queries

import (
	"github.com/jmoiron/sqlx"
	"github.com/monesonn/records.web/app/models"
)

type ClientQueries struct {
	*sqlx.DB
}

func (q *ClientQueries) GetClients() ([]models.Client, error) {
	// Define review variable.
	clients := []models.Client{}

	query := `SELECT * FROM client`

	// Send query to database.
	err := q.Select(&clients, query)
	if err != nil {
		// Return empty object and error.
		return clients, err
	}

	// Return query result.
	return clients, nil
}
