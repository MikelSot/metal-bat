package model

import (
	"encoding/json"
	"time"
)

type Service struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	IsActive    bool            `json:"is_active"`
	Description string          `json:"description"`
	Picture     json.RawMessage `json:"picture"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
