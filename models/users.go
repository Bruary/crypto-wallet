package models

import "time"

type User struct {
	ID        string
	UUID      string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Email     string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int
	Email     string
}
