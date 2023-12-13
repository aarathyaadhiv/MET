package useCaseInterface



type UserUseCase interface{
	SendOtp(phNo string) error 
}