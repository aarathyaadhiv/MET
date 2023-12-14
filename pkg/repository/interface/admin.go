package interfaces

import (
	"github.com/aarathyaadhiv/met/pkg/domain"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
)


type AdminRepository interface{
	IsAdminExist(email string)bool
	Save(admin models.Admin)(uint,error)
	FetchAdmin(email string)(domain.Admin,error)
}