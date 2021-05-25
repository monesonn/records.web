// Package models is used for represent data from database.
package models

import (
	"database/sql"
)

type Genre struct {
	GenreID int    `db:"genre_id" json:"id"`
	Name    string `db:"name" json:"name"`
}

type Artist struct {
	ArtistID int    `db:"artist_id" json:"id"`
	Name     string `db:"name" json:"name"`
}

type Record struct {
	RecordID    int            `db:"record_id" json:"record_id"`
	ArtistID    int            `db:"artist_id" json:"artist_id"`
	Title       string         `db:"title" json:"title"`
	Year        sql.NullInt32  `db:"year_" json:"year"`
	Country     sql.NullString `db:"country" json:"country"`
	GenreID     int            `db:"genre_id" json:"genre_id"`
	LabelID     sql.NullInt32  `db:"label_id" json:"label_id"`
	Description sql.NullString `db:"description" json:"description"`
}

type Product struct {
	ProductID  int    `db:"product_id" json:"product_id"`
	RecordID   int    `db:"record_id" json:"record_id"`
	MediumType string `db:"medium_type" json:"type"`
	Price      int    `db:"price" json:"price"`
	Quantity   int    `db:"quantity" json:"quantity"`
}
