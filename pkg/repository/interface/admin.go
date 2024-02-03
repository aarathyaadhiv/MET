package interfaces

import (
	"github.com/aarathyaadhiv/met/pkg/domain"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)


type AdminRepository interface{
	IsAdminExist(email string)bool
	Save(admin models.Admin)(uint,error)
	FetchAdmin(email string)(domain.Admin,error)
	IsUserExist(id uint)(bool,error)
	BlockUser(id uint)(uint,error)
	UnblockUser(id uint)(uint,error)
	GetUsers(page,count int)([]response.User,error)
	IsUserBlocked(id uint)(bool,error)
	GetSingleUser(id uint) (response.UserDetailsToAdmin, error)
	ReportedUsers()([]response.ReportedUsers,error)
	IsReportedUser(id uint)(bool,error)
	ReportedUser(reportId uint)(response.ReportedUser,error)
}