package models

import "time"

type Subscription struct {
	Name    string  `json:"name"`
	Amount  float64 `json:"amount"`
	Days    int     `json:"days"`
	Likes   int     `json:"like"`
	SeeLike bool    `json:"see_like" `
}

type Order struct {
	SubscriptionId uint      `json:"subscription_id"`
	UserId         uint      `json:"user_id"`
	SubscribeDate  time.Time `json:"subscribe_date"`
	Status         string    `json:"status"`
}

type OrderDetails struct {
	UserName string  `json:"user_name"`
	Amount   float64 `json:"amount"`
}

type PaymentRes struct {
	UserId         uint `json:"user_id"`
	SubscriptionId uint `json:"subscription_id"`
}
