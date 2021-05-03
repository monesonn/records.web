package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

func PostgreSQLConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to dbect to database: %v\n", err)
		os.Exit(1)
	}
	//defer db.Close()

	return db, err
}
