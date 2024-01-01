package response


import "time"



type Token struct {
	AccessToken  string
	RefreshToken string
}

type Id struct {
	Id uint
}

type Profile struct {
	UserDetails UserDetails
	Image       []string
	Interests   []string
}

type UserDetails struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Dob       time.Time `json:"dob"`
	Age       int       `json:"age"`
	PhNo      string    `json:"ph_no"`
	Gender    string    `json:"gender"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	Longitude string    `json:"longitude"`
	Lattitude string    `json:"lattitude"`
	Bio       string    `json:"bio"`
}