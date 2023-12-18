package interfaces

import (
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)




type UserRepository interface{
	IsUserExist(phNo string)(bool,error)
	IsUserBlocked(phNo string)(bool,error)
	FindByPhone(phNo string)(uint,error)
	CreateUserId(phNo string)(uint,error)
	UpdateUser(id uint,profile models.ProfileSave)(uint,error)
	AddInterest(id,interest uint)error
	AddImage(id uint,image string)error
	ShowProfile(id uint)(response.UserDetails,error)
	FetchImages(id uint)([]string,error)
	FetchInterests(id uint)([]string,error)
}