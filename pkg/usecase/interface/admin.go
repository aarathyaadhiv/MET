package useCaseInterface

import (
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)


type AdminUseCase interface{
	AdminSignUp(admin models.Admin) (uint, error)
	AdminLogin(admin models.Admin) (string, error)
	BlockUser(id uint)(uint,error)
	UnBlockUser(id uint)(uint,error)
	GetUsers()([]response.User,error)
}