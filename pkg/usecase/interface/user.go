package useCaseInterface

import (
	"mime/multipart"

	"github.com/aarathyaadhiv/met/pkg/domain"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)

type UserUseCase interface {
	SendOtp(phNo string) (response.SendOtp, error)
	VerifyOtp(otp models.OtpVerify) (bool, response.Token, error)
	AddProfile(profile models.Profile, id uint) (uint, error)
	ShowProfile(id uint) (response.Profile, error)
	UpdateUser(user models.UpdateUser, id uint) (response.Id, error)
	UpdateImage(id uint, image *multipart.Form) (response.Id, error)
	UpdatePreference(id uint, preference models.Preference) (response.Id, error)
	GetPreference(id uint) (models.Preference, error)
	Interests(id uint, user bool) ([]domain.Interests, error)
	Gender() ([]domain.Gender, error)
	DeleteAccount(userId uint) (response.Id, error)
	VerifyOTPtoUpdatePhNo(otp models.OtpVerify, userId uint) (response.SendOtp, error)
}
