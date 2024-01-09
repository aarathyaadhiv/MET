package response

import "time"

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Age     uint   `json:"age"`
	PhNo    string `json:"ph_no"`
	Gender  string `json:"gender"`
	City    string `json:"city"`
	Country string `json:"country"`
	IsBlock bool   `json:"is_block"`
}

type UserDetailsToAdmin struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Dob       time.Time `json:"dob"`
	Age       int       `json:"age"`
	PhNo      string    `json:"ph_no"`
	Gender    string    `json:"gender"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	Longitude float64   `json:"longitude"`
	Lattitude float64   `json:"lattitude"`
	Bio       string    `json:"bio"`
	Images    string    `json:"image"`
	Interests string    `json:"interests"`
}
