package domain

import "time"

type User struct {
	Id             uint         `json:"id" gorm:"primaryKey"`
	Name           string       `json:"name"`
	Dob            time.Time    `json:"dob"`
	Age            uint         `json:"age"`
	PhNo           string       `json:"ph_no"`
	GenderId       uint         `json:"gender_id"`
	Gender         Gender       `json:"gender" gorm:"foreignKey:GenderId"`
	City           string       `json:"city"`
	Country        string       `json:"country"`
	Longitude      float64      `json:"longitude"`
	Lattitude      float64      `json:"lattitude"`
	Bio            string       `json:"bio"`
	IsBlock        bool         `json:"is_block" gorm:"default:false"`
	LikeCount      int          `json:"like_count" gorm:"default:5"`
	IsSubscribed   bool         `json:"is_subscribed" gorm:"default:false"`
	SubscriptionId uint         `json:"subscription_id"`
	Subscription   Subscription `json:"subscription" gorm:"foreignKey:SubscriptionId"`
}

type Gender struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" `
}

type UserInterests struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	UserId     uint      `json:"user_id"`
	User       User      `json:"user" gorm:"foreignKey:UserId" `
	InterestId uint      `json:"interest_id"`
	Interest   Interests `json:"interest" gorm:"foreignKey:InterestId"`
}

type Interests struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Interest string `json:"interest"`
}

type Images struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	UserId uint   `json:"user_id"`
	User   User   `json:"user" gorm:"foreignKey:UserId" `
	Image  string `json:"image" `
}

type Preference struct {
	Id          uint `json:"id" gorm:"primaryKey"`
	UserId      uint `json:"user_id"`
	User        User `json:"user" gorm:"foreignKey:UserId" `
	MinAge      uint `json:"min_age"`
	MaxAge      uint `json:"max_age"`
	Gender      uint `json:"gender"`
	MaxDistance uint `json:"max_distance"`
}
