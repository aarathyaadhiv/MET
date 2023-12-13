package models



type OtpRequest struct{
	PhNo string `json:"ph_no" binding:"required" validate:"required"`
}

type OtpVerify struct{
	PhNo string `json:"ph_no" binding:"required" validate:"required"`
	Code string `json:"code" binding:"required" validate:"required"`
}