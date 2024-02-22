package domain

import "time"

type Admin struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Subscription struct {
	Id           uint    `json:"id" gorm:"primaryKey"`
	Name         string  `json:"name"`
	Amount       float64 `json:"amount"`
	Days         int     `json:"days"`
	Likes        int     `json:"likes"`
	RewindCount  int     `json:"rewind_count"`
	HideAdds     bool    `json:"hide_adds"`
	PriorityLike bool    `json:"priority_like"`
	SeeLike      bool    `json:"see_like"`
	IsActive     bool    `json:"is_active" gorm:"default:true"`
}

type SubscriptionOrder struct {
	Id             uint         `json:"id" gorm:"primaryKey"`
	SubscriptionId uint         `json:"subscription_id"`
	Subscription   Subscription `json:"subscription" gorm:"foreignKey:SubscriptionId;constraint:OnDelete:CASCADE"`
	UserId         uint         `json:"user_id"`
	User           User         `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	SubscribeDate  time.Time    `json:"subscribe_date"`
	Status         string       `json:"status"`
	PaymentId      string       `json:"payment_id"`
	RazorId        string       `json:"razor_id"`
}
