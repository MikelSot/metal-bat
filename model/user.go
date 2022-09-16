package model

import (
	"regexp"
	"strings"
	"time"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

const (
	lenPassword = 6
	minLenEmail = 6
	maxLenEmail = 254
)

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

func (u User) HasID() bool { return u.ID > 0 }

func (u User) IsValidLenPassword() bool { return len(u.Password) >= lenPassword }

func (u User) IsStringEmpty(data string) bool { return strings.TrimSpace(data) == "" }

func (u User) IsEmailValidByRegex() bool { return emailRegex.MatchString(u.Email) }

func (u User) IsValidLenEmail() bool {
	return len(u.Email) >= minLenEmail && len(u.Email) <= maxLenEmail
}

type Users []User
