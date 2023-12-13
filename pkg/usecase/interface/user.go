package useCaseInterface

import "github.com/aarathyaadhiv/met/pkg/utils/models"



type UserUseCase interface{
	SendOtp(phNo string) error 
	VerifyOtp(otp models.OtpVerify)(bool,string,error)
}