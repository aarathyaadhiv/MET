package interfaces

import "github.com/aarathyaadhiv/met/pkg/utils/models"




type UserRepository interface{
	IsUserExist(phNo string)(bool,error)
	IsUserBlocked(phNo string)(bool,error)
	FindByPhone(phNo string)(uint,error)
	CreateUserId(phNo string)(uint,error)
	UpdateUser(id uint,profile models.ProfileSave)(uint,error)
	AddInterest(id,interest uint)error
	AddImage(id uint,image string)error
}