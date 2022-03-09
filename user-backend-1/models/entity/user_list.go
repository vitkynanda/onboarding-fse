package entity

import "github.com/google/uuid"

type UserList struct {
	Id         uuid.UUID `json:"id"`
	Firstname  string    `json:"firstName" binding:"required"`
	Lastname   string    `json:"lastName"  binding:"required"`
	Gender     string    `json:"gender" binding:"required"`
	Birthdate  string    `json:"birthdate" binding:"required"`
	Active     bool      `json:"active" binding:"required"`
}