package domain

import "time"

type Admin struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Subscription struct {
	Id     uint    `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Days   int     `json:"days"`
}

type Subscription_order struct {
	Id             uint         `json:"id" gorm:"primaryKey"`
	SubscriptionId uint         `json:"subscription_id"`
	Subscription   Subscription `json:"subscription" gorm:"foreignKey:SubscriptionId"`
	UserId         uint         `json:"user_id"`
	User           User         `json:"user" gorm:"foreignKey:UserId"`
	SubscribeDate  time.Time    `json:"subscribe_date"`
	Status         string       `json:"status"`
	PaymentId      string       `json:"payment_id"`
	RazorId        string       `json:"razor_id"`
}
