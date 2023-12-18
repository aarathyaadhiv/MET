package models

import (
	"mime/multipart"
	"time"
)



type OtpRequest struct{
	PhNo string `json:"ph_no" binding:"required" validate:"required"`
}

type OtpVerify struct{
	PhNo string `json:"ph_no" binding:"required" validate:"required"`
	Code string `json:"code" binding:"required" validate:"required"`
}

type Profile struct{
	Name string `json:"name" binding:"required" validate:"required"`
	Dob time.Time `json:"-" binding:"required" validate:"required" `
	GenderId uint `json:"gender_id" binding:"required" validate:"required"`
	City string `json:"city" binding:"required" validate:"required"`
	Country string `json:"country" binding:"required" validate:"required"`
	Longitude string `json:"longitude" binding:"required" validate:"required"`
	Lattitude string `json:"lattitude" binding:"required" validate:"required"`
	Bio string `json:"bio" binding:"required" validate:"required"`
	Image *multipart.Form `json:"image" binding:"required" validate:"required"`
	Interests []uint `json:"interests" binding:"required" validate:"required"`
}

type ProfileSave struct{
	Name string `json:"name" binding:"required" validate:"required"`
	Dob time.Time `json:"-" binding:"required" validate:"required" `
	Age int `json:"-" binding:"required" validate:"required" `
	GenderId uint `json:"gender_id" binding:"required" validate:"required"`
	City string `json:"city" binding:"required" validate:"required"`
	Country string `json:"country" binding:"required" validate:"required"`
	Longitude string `json:"longitude" binding:"required" validate:"required"`
	Lattitude string `json:"lattitude" binding:"required" validate:"required"`
	Bio string `json:"bio" binding:"required" validate:"required"`
}