package domain

import "time"

type Likes struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE" `
	LikedId   uint      `json:"liked_id"`
	LikedUser User      `json:"liked_user" gorm:"foreignKey:LikedId;constraint:OnDelete:CASCADE"`
	Time      time.Time `json:"time"`
}

type Match struct {
	Id          uint `json:"id" gorm:"primaryKey"`
	UserId      uint `json:"user_id"`
	User        User `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE" `
	MatchId     uint `json:"match_id"`
	MatchedUser User `json:"matched_user" gorm:"foreignKey:MatchId;constraint:OnDelete:CASCADE"`
	Time      time.Time `json:"time"`
}

type BlockedUsers struct {
	Id          uint `json:"id" gorm:"primaryKey"`
	UserId      uint `json:"user_id"`
	User        User `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE" `
	BlockedId   uint `json:"blocked_id"`
	BlockedUser User `json:"blocked_user" gorm:"foreignKey:BlockedId;constraint:OnDelete:CASCADE"`
}

type ReportedUsers struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	UserId       uint   `json:"user_id"`
	User         User   `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE" `
	ReportedId   uint   `json:"reported_id"`
	ReportedUser User   `json:"blocked_user" gorm:"foreignKey:ReportedId;constraint:OnDelete:CASCADE"`
	Message      string `json:"message"`
	Time      time.Time `json:"time"`
}
