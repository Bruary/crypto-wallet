package models

import "time"

type User struct {
	ID        string
	FirstName string
	LastName  string
	FullName  string
	Age       string
	DOB       time.Time
	Email     string
}
