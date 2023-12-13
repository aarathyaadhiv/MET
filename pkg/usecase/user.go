package usecase

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/aarathyaadhiv/met/pkg/helper"
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
)

type UserUseCase struct {
	Repo   interfaces.UserRepository
	Config config.Config
}

func NewUserUseCase(repo interfaces.UserRepository, config config.Config) useCaseInterface.UserUseCase {
	return &UserUseCase{Repo: repo,
		Config: config}
}

func (u *UserUseCase) SendOtp(phNo string) error {
	
	regex := regexp.MustCompile(`^\d{10}$`)
	if !regex.MatchString(phNo){
		return errors.New("the given phone number is not in the correct format")
	}
	exist,err:=u.Repo.IsUserExist(phNo)
	if err!=nil{
		return err
	}
	if exist{
		block,err:=u.Repo.IsUserBlocked(phNo)
		if err!=nil{
			return err
		}
		if block{
			return errors.New("this number is blocked for this app")
		}
	}
	phone:="+91"+phNo
	helper.TwillioSetup(u.Config.TwilioAccountSID,u.Config.TwilioAuthToken)
	_,err=helper.SendOtp(phone,u.Config.TwilioServicesId)	
	if err!=nil{
		return errors.New("error in sending otp")
	}
	return nil
}

func (u *UserUseCase) VerifyOtp(otp models.OtpVerify)(bool,string,error){
	regex := regexp.MustCompile(`^\d{10}$`)
	if !regex.MatchString(otp.PhNo){
		return false,"",errors.New("the given phone number is not in the correct format")
	}
	helper.TwillioSetup(u.Config.TwilioAccountSID,u.Config.TwilioAuthToken)
	err:=helper.ValidateOtp(otp,u.Config.TwilioServicesId)
	if err!=nil{
		return false,"",errors.New("otp verification failed")
	}
	//user search
	exist,err:=u.Repo.IsUserExist(otp.PhNo)
	if err!=nil{
		return false,"",err
	}
	var id uint
	if exist{
		id,err=u.Repo.FindByPhone(otp.PhNo)
		if err!=nil{
			return false,"",err
		}
	}else{
		id,err=u.Repo.CreateUserId(otp.PhNo)
		if err!=nil{
			return false,"",err
		}
	}
	
	fmt.Print(id)
	//generate token
	token,err:=helper.GenerateUserToken(id)
	if err!=nil{
		return false,"",err
	}
	return exist,token,nil
}
