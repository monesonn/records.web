package models

import "time"

type Client struct {
	ID        string    `db:"id" json:"id"`
	Username  string    `db:"username" json:"username" validate:"required,lte=255"`
	Email     string    `db:"email" json:"email" validate:"required,email,lte=255"`
	FirstName string    `db:"first_name" json:"fname"`
	LastName  string    `db:"last_name" json:"lname"`
	Telno     string    `db:"tel_no" json:"telno"`
	Gender    string    `db:"gender" json:"gender"`
	Birthday  string    `db:"birthday" json:"birthday"`
	Picture   string    `db:"picture" json:"picture"`
	CreatedAt time.Time `db:"inserted_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Profile struct {
	ID        string `db:"user_id" json:"uuid"`
	FirstName string `db:"first_name" json:"fname"`
	LastName  string `db:"last_name" json:"lname"`
	Telno     string `db:"tel_no" json:"telno"`
	Gender    string `db:"gender" json:"gender"`
	Birthday  string `db:"birthday" json:"birthday"`
	Picture   int    `db:"picture" json:"picture"`
}

type Review struct {
	UserID    string `db:"user_id" json:"user_id"`
	ProductID int    `db:"product_id" json:"product_id"`
	ViewRate  int    `db:"rate" json:"rate"`
	ViewDate  string `db:"view_date" json:"date"`
	Comment   string `db:"commentary" json:"comment"`
}
