package model

import "time"

type Category struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	IsActive    bool      `json:"is_active"`
	IsPromotion bool      `json:"is_promotion"`
	Picture     string    `json:"picture"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
