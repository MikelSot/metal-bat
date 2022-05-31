package model

import "time"

type User struct {
	ID                 uint           `json:"id"`
	Firstname          string         `json:"firstname"`
	Lastname           string         `json:"lastname"`
	Nickname           string         `json:"nickname"`
	Email              string         `json:"email"`
	Password           string         `json:"password"`
	IdentificationType Identification `json:"identification_type"`
	PhoneNumber        string         `json:"phone_number"`
	Picture            string         `json:"picture"`
	Address            string         `json:"address"`
	District           string         `json:"district"`
	Instagram          string         `json:"instagram"`
	Twitter            string         `json:"twitter"`
	Facebook           string         `json:"facebook"`
	Gender             string         `json:"gender"`
	Biography          string         `json:"biography"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          time.Time      `json:"deleted_at"`
}

type Users []User
