package database

import (
	"github.com/monesonn/records.web/app/queries"
)

type Queries struct {
	*queries.UserQueries    // load queries from User model
	*queries.RecordsQueries // load queries from Records model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries:    &queries.UserQueries{DB: db},    // from User model
		RecordsQueries: &queries.RecordsQueries{DB: db}, // from Records model
	}, nil
}
