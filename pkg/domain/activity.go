package domain

import "time"

type Likes struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:UserId" `
	LikedId   uint      `json:"liked_id"`
	LikedUser User      `json:"liked_user" gorm:"foreignKey:LikedId"`
	Time      time.Time `json:"time"`
}

type Match struct {
	Id          uint `json:"id" gorm:"primaryKey"`
	UserId      uint `json:"user_id"`
	User        User `json:"user" gorm:"foreignKey:UserId" `
	MatchId     uint `json:"match_id"`
	MatchedUser User `json:"matched_user" gorm:"foreignKey:MatchId"`
	Time      time.Time `json:"time"`
}

type BlockedUsers struct {
	Id          uint `json:"id" gorm:"primaryKey"`
	UserId      uint `json:"user_id"`
	User        User `json:"user" gorm:"foreignKey:UserId" `
	BlockedId   uint `json:"blocked_id"`
	BlockedUser User `json:"blocked_user" gorm:"foreignKey:BlockedId"`
}

type ReportedUsers struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	UserId       uint   `json:"user_id"`
	User         User   `json:"user" gorm:"foreignKey:UserId" `
	ReportedId   uint   `json:"reported_id"`
	ReportedUser User   `json:"blocked_user" gorm:"foreignKey:ReportedId"`
	Message      string `json:"message"`
	Time      time.Time `json:"time"`
}
