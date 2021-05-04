package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type User struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Email        string    `db:"login" json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `db:"password" json:"password_hash,omitempty" validate:"required,lte=255"`
	UserStatus   bool      `db:"active" json:"user_status"`
	UserRole     string    `db:"category" json:"user_role" validate:"required,lte=25"`
	CreatedAt    time.Time `db:"inserted_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
