package dto

import "github.com/google/uuid"


type User struct {
	Id         uuid.UUID `json:"id"`
	Firstname  string    `json:"firstName" binding:"required"`
	Lastname   string    `json:"lastName"  binding:"required"`
	Gender     string    `json:"gender" binding:"required"`
	Birthdate  string    `json:"birthdate" binding:"required"`
	Active     bool      `json:"active" binding:"required"`
	Contact    Contact 	 `json:"contact" binding:"required"`
	Hobbies    []string  `json:"hobbies" binding:"required"`
}