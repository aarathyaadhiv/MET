package useCaseInterface

import "github.com/aarathyaadhiv/met/pkg/utils/models"


type AdminUseCase interface{
	AdminSignUp(admin models.Admin) (uint, error)
	AdminLogin(admin models.Admin) (string, error)
}