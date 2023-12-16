package domain


type User struct{
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Dob string `json:"dob"`
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

type Interests struct{
	Id uint `json:"id" gorm:"primaryKey"`
	UserId uint `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:UserId" `
	Interest string `json:"interest" `
}

type Images struct{
	Id uint `json:"id" gorm:"primaryKey"`
	UserId uint `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:UserId" `
	Image string `json:"image" `
}