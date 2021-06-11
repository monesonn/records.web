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
	OrderStatus string    `db:"order_status" json:"status"`
	LastUpdate  time.Time `db:"last_update" json:"last_update"`
}

type OrderList struct {
	OrderItemID     int    `db:"order_item_id" json:"item_id"`
	ProductID       int    `db:"product_id" json:"product_id"`
	OrderID         int    `db:"order_id" json:"order_id"`
	OrderItemQty    int    `db:"order_item_quantity" json:"item_qty"`
	OrderItemCost   int    `db:"order_item_price" json:"item_cost"`
	OrderItemStatus string `db:"order_item_status" json:"item_status"`
}

// type OrderItems struct
