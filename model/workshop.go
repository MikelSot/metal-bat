package model

import (
	"encoding/json"
	"time"
)

type Workshop struct {
	ID          uint16          `json:"id"`
	Name        string          `json:"name"`
	BossID      uint            `json:"boss_id"`
	Pictures    json.RawMessage `json:"pictures"`
	Address     string          `json:"address"`
	Description string          `json:"description"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   time.Time       `json:"deleted_at"`
}
