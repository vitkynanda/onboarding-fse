package dto

type Contact struct {
	Email string `json:"email" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}