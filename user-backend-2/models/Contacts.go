package models

type Contacts struct {
	Email string `json:"email" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}