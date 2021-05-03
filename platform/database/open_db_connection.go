package database

import (
	"github.com/monesonn/records.web/app/queries"
)

type Queries struct {
	*queries.RecordsQueries // load queries from model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{&queries.RecordsQueries{DB: db}}, nil
}
