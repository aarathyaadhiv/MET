package repository

import (
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	"gorm.io/gorm"
)


type UserRepository struct{
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB)interfaces.UserRepository{
	return &UserRepository{db}
}

func (u *UserRepository) IsUserExist(phNo string)(bool,error){
	var count int
	if err:=u.DB.Raw(`SELECT COUNT(*) FROM users WHERE ph_no=? `,phNo).Scan(&count).Error;err!=nil{
		return false,err
	}
	return count>0,nil
}

func (u *UserRepository) IsUserBlocked(phNo string)(bool,error){
	var block bool
	if err:=u.DB.Raw(`SELECT is_block FROM users WHERE ph_no=?`,phNo).Scan(&block).Error;err!=nil{
		return false,err
	}
	return block,nil
}

func (u *UserRepository) FindByPhone(phNo string)(uint,error){
	var id uint
	if err:=u.DB.Raw(`SELECT id FROM users WHERE ph_no=?`,phNo).Scan(&id).Error;err!=nil{
		return 0,err
	}
	return id,nil
}

func (u *UserRepository) CreateUserId(phNo string)(uint,error){
	var id uint
	if err:=u.DB.Raw(`INSERT INTO users(ph_no) VALUES(?) RETURNING id `,phNo).Scan(&id).Error;err!=nil{
		return 0,err
	}
	return id,nil
}