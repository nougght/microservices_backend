package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
	// Description string `json:"description"`
	// Image        string    `db:"image_url" json:"image"`
	IsActive     bool      `db:"is_active" json:"is_active"`
	CreationDate time.Time `db:"created_at" json:"created_at"`
	UpdateDate   time.Time `db:"updated_at" json:"updated_at"`
}
