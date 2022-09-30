package model

import "time"

type UserRole struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	RoleID    uint16    `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
