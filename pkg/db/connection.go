package db

import (
	"fmt"

	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/aarathyaadhiv/met/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(c config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", c.DBHost, c.DBUser, c.DBName, c.DBPort, c.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.Interests{})
	db.AutoMigrate(&domain.UserInterests{})
	// err := db.Exec(domain.UserInterests{}.Migration()).Error
	// if err != nil {
	// 	return db, err
	// }

	db.AutoMigrate(&domain.Images{})
	// err := db.Exec(domain.Images{}.Migration()).Error
	// if err != nil {
	// 	return db, err
	// }
	db.AutoMigrate(&domain.Gender{})
	db.AutoMigrate(&domain.Subscription{})
	db.AutoMigrate(&domain.SubscriptionOrder{})
	db.AutoMigrate(&domain.Likes{})
	db.AutoMigrate(&domain.Match{})
	db.AutoMigrate(&domain.BlockedUsers{})
	db.AutoMigrate(&domain.ReportedUsers{})
	db.AutoMigrate(&domain.Preference{})

	return db, dbErr
}
