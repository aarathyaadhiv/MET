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
