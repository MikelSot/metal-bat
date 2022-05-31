package model

import "time"

type Employee struct {
	ID              uint      `json:"id"`
	UserID          uint      `json:"user_id"`
	Birthday        time.Time `json:"birthday"`
	IsActive        bool      `json:"is_active"`
	Salary          float32   `json:"salary"`
	ShiftType       Shift     `json:"shift_type"`
	Profession      string    `json:"profession"`
	PositionType    string    `json:"position_type"`
	EmergencyNumber string    `json:"emergency_number"`
	WorkshopID      uint16    `json:"workshop_id"`
	StartDate       time.Time `json:"start_date"`
	FinishDate      time.Time `json:"finish_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}
