package interfaces

import (
	"github.com/aarathyaadhiv/met/pkg/domain"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)

type UserRepository interface {
	IsUserExist(phNo string) (bool, error)
	IsUserBlocked(phNo string) (bool, error)
	FindByPhone(phNo string) (uint, error)
	CreateUserId(phNo string) (uint, error)
	UpdateUser(id uint, profile models.ProfileSave) (uint, error)
	IsInterestExist(id, interest uint) (bool, error)
	AddInterest(id, interest uint) error
	DeleteInterest(id uint) error
	IsImageExist(id uint, image string) (bool, error)
	AddImage(id uint, image string) error
	DeleteImage(id uint) error
	ShowProfile(id uint) (response.UserDetails, error)
	FetchImages(id uint) ([]string, error)
	FetchInterests(id uint) ([]string, error)
	IsBlocked(id uint) (bool, error)
	UpdateLocation(id uint, location models.UpdateLocation) (uint, error)
	UpdatePhNo(id uint,phNo string)error
	UpdateUserDetails(id uint, user models.UpdateUserDetails) error
	AddPreference(id uint, preference models.Preference) error
	UpdatePreference(id uint, preference models.Preference) (uint, error)
	GetPreference(id uint) (models.Preference, error)
	FetchShortDetail(id uint) (models.UserShortDetail, error)
	ShowInterests(id uint) ([]domain.Interests, error)
	Interests() ([]domain.Interests, error)
	Gender() ([]domain.Gender, error)
	DeleteUser(userId uint)(uint,error)
}
