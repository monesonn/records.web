package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type User struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Email        string    `db:"email" json:"email" validate:"required,email,lte=255"`
	Username     string    `db:"username" json:"username" validate:"required,lte=255"`
	PasswordHash string    `db:"password" json:"password_hash,omitempty" validate:"required,lte=255"`
	UserRole     string    `db:"category" json:"category" validate:"required,lte=25"`
	UserStatus   bool      `db:"active" json:"user_status"`
	CreatedAt    time.Time `db:"inserted_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
