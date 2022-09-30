package model

import "time"

type Appointment struct {
	ID              uint              `json:"id"`
	UserID          uint              `json:"user_id"`
	WorkshopID      uint16            `json:"workshop_id"`
	ServiceID       uint              `json:"service_id"`
	VehicleID       uint              `json:"vehicle_id"`
	Description     string            `json:"description"`
	AppointmentDate time.Time         `json:"appointment_date"`
	DeliveryDate    time.Time         `json:"delivery_date"`
	AttentionOrder  uint16            `json:"attention_order"`
	State           AppointmentStatus `json:"state"`
	PickUp          bool              `json:"pick_up"`
	Delivery        bool              `json:"delivery"`
	IsCanceled      bool              `json:"is_canceled"`
	Slug            string            `json:"slug"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       time.Time         `json:"deleted_at"`
}

type Appointments []Appointment
