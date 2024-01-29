package response

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
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	Dob       string  `json:"dob"`
	Age       int     `json:"age"`
	PhNo      string  `json:"ph_no"`
	Gender    string  `json:"gender"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Longitude float64 `json:"longitude"`
	Lattitude float64 `json:"lattitude"`
	Bio       string  `json:"bio"`
}

type Home struct {
	Id        uint     `json:"id"`
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Gender    string   `json:"gender"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	Longitude float64  `json:"longitude"`
	Lattitude float64  `json:"lattitude"`
	Bio       string   `json:"bio"`
	Images    []string `json:"images"`
}
