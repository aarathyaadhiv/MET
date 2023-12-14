package repository

import (
	"github.com/aarathyaadhiv/met/pkg/domain"
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"gorm.io/gorm"
)


type AdminRepository struct{
	DB *gorm.DB
}


func NewAdminRepository(db *gorm.DB)interfaces.AdminRepository{
	return &AdminRepository{db}
}

func (a *AdminRepository) IsAdminExist(email string)bool{
	var count int
	if err:=a.DB.Raw(`SELECT COUNT(*) FROM admins WHERE email=?`,email).Scan(&count).Error;err!=nil{
		return false
	}
	return count>0
}
 
func (a *AdminRepository) Save(admin models.Admin)(uint,error){
	var id uint
	if err:=a.DB.Raw(`INSERT INTO admins(email,password) VALUES(?,?) RETURNING id`,admin.Email,admin.Password).Scan(&id).Error;err!=nil{
		return 0,err
	}
	return id,nil
}

func (a *AdminRepository) FetchAdmin(email string)(domain.Admin,error){
	var admin domain.Admin
	if err:=a.DB.Raw(`SELECT * FROM admins WHERE email=?`,email).Scan(&admin).Error;err!=nil{
		return domain.Admin{},err
	}
	return admin,nil
}