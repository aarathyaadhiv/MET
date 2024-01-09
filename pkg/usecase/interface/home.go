package useCaseInterface

import "github.com/aarathyaadhiv/met/pkg/utils/response"


type HomeUseCase interface{
	HomePage(id uint,page,count int) ([]response.Home, error)
}