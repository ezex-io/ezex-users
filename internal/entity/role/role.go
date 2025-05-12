package role

import (
	"time"
)

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	CreatedByID string    `json:"created_by_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedByID string    `json:"updated_by_id"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedByID string    `json:"deleted_by_id"`
	DeletedAt   time.Time `json:"deleted_at"`

	IsSystem  bool `json:"is_system"`
	IsDefault bool `json:"is_default"`
}
