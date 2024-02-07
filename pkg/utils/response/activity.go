package response

import "time"

type Like struct {
	UserId  uint `json:"user_id"`
	LikedId uint `json:"liked_id"`
}
type ShowUserDetails struct {
	Id      uint      `json:"id"`
	Name    string    `json:"name"`
	Dob     time.Time `json:"dob"`
	Age     int       `json:"age"`
	Gender  string    `json:"gender"`
	City    string    `json:"city"`
	Country string    `json:"country"`
	Bio     string    `json:"bio"`
	Image   string    `json:"image"`
}
type ShowLike struct {
	UserId       uint              `json:"user_id"`
	IsSubscribed bool              `json:"is_subscribed"`
	SeeLike      bool              `json:"see_like"`
	Likes        []ShowUserDetails `json:"likes"`
}

type UnMatch struct {
	UserId    uint `json:"user_id"`
	MatchedId uint `json:"matched_id"`
}

type ShowMatch struct {
	UserId  uint              `json:"user_id"`
	Matches []ShowUserDetails `json:"matches"`
}

type BlockUser struct {
	UserId    uint `json:"user_id"`
	BlockedId uint `json:"blocked_id"`
}

type Report struct {
	UserId     uint `json:"user_id"`
	ReportedId uint `json:"reported_id"`
}
