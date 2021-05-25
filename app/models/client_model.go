package models

import (
	"database/sql"
)

type Client struct {
	ClientID  int            `db:"client_id" json:"id"`
	FirstName string         `db:"first_name" json:"fname"`
	LastName  string         `db:"last_name" json:"lname"`
	Telno     string         `db:"tel_no" json:"telno"`
	Gender    string         `db:"gender" json:"gender"`
	Birthday  sql.NullString `db:"birthday" json:"birthday"`
	Picture   sql.NullInt32  `db:"picture" json:"picture"`
}
