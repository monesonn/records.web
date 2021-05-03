// Package models is used for represent data from database.
package models

type Genre struct {
	GenreID int32  `db:"genre_id" json:"id"`
	Name    string `db:"name" json:"name"`
}

type Artist struct {
	ArtistID int32  `db:"artist_id" json:"id"`
	Name     string `db:"name" json:"name"`
}
