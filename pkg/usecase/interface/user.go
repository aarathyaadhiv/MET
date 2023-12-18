package useCaseInterface

import (
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)



type UserUseCase interface{
	SendOtp(phNo string) error 
	VerifyOtp(otp models.OtpVerify)(bool,response.Token,error)
	AddProfile(profile models.Profile,id uint)(uint,error)
}