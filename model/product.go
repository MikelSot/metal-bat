package model

import "time"

type Product struct {
	ID          uint      `json:"id"`
	Name        uint      `json:"name"`
	Stock       uint16    `json:"stock"`
	Slug        string    `json:"slug"`
	Picture     string    `json:"picture"`
	Description string    `json:"description"`
	IsPromotion bool      `json:"is_promotion"`
	Video       string    `json:"video"`
	CategoryID  uint      `json:"category_id"`
	WorkshopID  uint16    `json:"workshop_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
