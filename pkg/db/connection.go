package db

import (
	"fmt"

	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/aarathyaadhiv/met/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func ConnectDB(c config.Config)(*gorm.DB,error){
	psqlInfo:=fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s",c.DBHost,c.DBUser,c.DBName,c.DBPort,c.DBPassword)
	db,dbErr:=gorm.Open(postgres.Open(psqlInfo),&gorm.Config{SkipDefaultTransaction: true})

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.Interests{})
	db.AutoMigrate(&domain.UserInterests{})
	db.AutoMigrate(&domain.Images{})
	db.AutoMigrate(&domain.Gender{})
	db.AutoMigrate(&domain.Subscription{})
	db.AutoMigrate(&domain.Subscription_order{})
	
	return db,dbErr
}