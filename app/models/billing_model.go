// Package models is used for represent data from database.
package models

import "time"

type Invoice struct {
	InvoiceID     int       `db:"invoice_id" json:"id"`
	OrderID       int       `db:"order_id" json:"order_id"`
	InvoiceStatus bool      `db:"invoice_status" json:"status"`
	LastUpdate    time.Time `db:"last_update" json:"last_update"`
}

type Orders struct {
	OrderID     int       `db:"order_id" json:"order_id"`
	UserID      string    `db:"user_id" json:"user_id"`
	TotalCost   int       `db:"total_cost" json:"total"`
	OrderStatus bool      `db:"order_status" json:"status"`
	LastUpdate  time.Time `db:"last_update" json:"last_update"`
}

// type OrderItems struct
