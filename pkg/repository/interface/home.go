package interfaces

import (
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)

type HomeRepository interface {
	FetchUser(id uint) (models.FetchUser, error)
	FetchPreference(id uint) (models.Preference, error)
	FetchUsers(maxAge, minAge int, gender, id uint) ([]response.Home, error)
	FetchImages(id uint) ([]string, error)
	FetchInterests(id uint) ([]string, error)
	FetchInterestId(id uint) ([]uint, error)
	IsLikeExist(userId, likedId uint) (bool, error)
	IsBlocked(userId, blockedId uint) (bool, error)
	FetchUserWithInterest(id uint, interestId []uint) ([]response.Home, error)
	IsInterestValid(id uint, interestId uint) (bool, error)
	FetchUserByInterest(id uint, interestId uint) ([]response.Home, error)
}
