package useCaseInterface

import (
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)

type HomeUseCase interface {
	HomePage(id uint, page, count int, interest bool, interestId string) ([]response.Home, error)
}
