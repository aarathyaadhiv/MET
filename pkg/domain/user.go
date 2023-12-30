package domain

import "time"


type User struct{
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Dob time.Time `json:"dob"`
	Age uint `json:"age"`
	PhNo string `json:"ph_no"`
	GenderId uint `json:"gender_id"`
	Gender Gender `json:"gender" gorm:"foreignKey:GenderId"`
	City string `json:"city"`
	Country string `json:"country"`
	Longitude string `json:"longitude"`
	Lattitude string `json:"lattitude"`
	Bio string `json:"bio"`
	IsBlock bool `json:"is_block" gorm:"default:false"`
	ReportCount int `json:"report_count" gorm:"default:0"`
}

type Gender struct{
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" `
}

type UserInterests struct{
	Id uint `json:"id" gorm:"primaryKey"`
	UserId uint `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:UserId" `
	InterestId uint `json:"interest_id"`
	Interest Interests `json:"interest" gorm:"foreignKey:InterestId"`
}

type Interests struct{
	Id uint `json:"id" gorm:"primaryKey"`
	Interest string `json:"interest"`
}

type Images struct{
	Id uint `json:"id" gorm:"primaryKey"`
	UserId uint `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:UserId" `
	Image string `json:"image" `
}

type Likes struct{
	Id uint `json:"id" gorm:"primaryKey"`
	UserId uint `json:"user_id"`
	LikedId uint `json:"liked_id"`
}

type Match struct{
	Id uint `json:"id" gorm:"primaryKey"`
	UserId uint `json:"user_id"`
	MatchId uint `json:"match_id"`
}