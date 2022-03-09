package models

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID
	Firstname string
	Lastname  string
	Gender    string
	Birthdate string
	Email     string
	Phone     string
	Active    bool
	Hobbies   string
}