package interfaces




type UserRepository interface{
	IsUserExist(phNo string)(bool,error)
	IsUserBlocked(phNo string)(bool,error)
	FindByPhone(phNo string)(uint,error)
	CreateUserId(phNo string)(uint,error)
}