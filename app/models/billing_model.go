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
	OrderID     int       `db:"order_id" json:"id"`
	SupplierID  int       `db:"supplier_id" json:"supplier_id"`
	ClientID    int       `db:"client_id" json:"client_id"`
	OrderStatus bool      `db:"order_status" json:"status"`
	LastUpdate  time.Time `db:"last_update" json:"last_update"`
}

// type OrderItems struct
