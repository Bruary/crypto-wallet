package models

import "time"

type User struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Age       string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
