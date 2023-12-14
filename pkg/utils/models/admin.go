package models


type Admin struct{
	Email string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required" validate:"required"`
}