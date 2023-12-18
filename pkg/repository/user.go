package repository

import (
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
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

func (u *UserRepository) UpdateUser(id uint,profile models.ProfileSave)(uint,error){
	var userId uint
	if err:=u.DB.Raw(`UPDATE users SET name=?,dob=?,age=?,gender_id=?,city=?,country=?,longitude=?,lattitude=?,bio=? WHERE id=? RETURNING id`,profile.Name,profile.Dob,profile.Age,profile.GenderId,profile.City,profile.Country,profile.Longitude,profile.Lattitude,profile.Bio,id).Scan(&userId).Error;err!=nil{
		return 0,err
	}
	return userId,nil
}

func (u *UserRepository) AddInterest(id,interest uint)error{
	if err:=u.DB.Exec(`INSERT INTO user_interests(user_id,interest_id) values(?,?)`,id,interest).Error;err!=nil{
		return err
	}
	return nil
}

func (u *UserRepository) AddImage(id uint,image string)error{
	if err:=u.DB.Exec(`INSERT INTO images(user_id,image) values(?,?)`,id,image).Error;err!=nil{
		return err
	}
	return nil
}