package queries

import (
	"github.com/jmoiron/sqlx"
	"github.com/monesonn/records.web/app/models"
)

type ClientQueries struct {
	*sqlx.DB
}

func (q *ClientQueries) GetClients() ([]models.Client, error) {
	clients := []models.Client{}
	query := `SELECT * FROM client`
	err := q.Select(&clients, query)
	if err != nil {
		return clients, err
	}
	return clients, nil
}

func (q *RecordsQueries) GetClient(id int) (models.Client, error) {
	client := models.Client{}
	query := `SELECT * FROM client where client_id = $1`
	err := q.Get(&client, query, id)
	if err != nil {
		return client, err
	}
	return client, nil
}
