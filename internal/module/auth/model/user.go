package model

import "time"

type UserStatus string

const (
	UserLocked UserStatus = "LOCKED"
	UserActive UserStatus = "LOCKED"
)

type User struct {
	ID            int64
	Username      string
	Email         string
	PhoneNumber   string
	FirstName     string
	LastName      *string
	Status        UserStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}