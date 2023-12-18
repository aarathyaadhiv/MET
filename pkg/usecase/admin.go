package usecase

import (
	"errors"

	"github.com/aarathyaadhiv/met/pkg/helper"
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	Repo interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) useCaseInterface.AdminUseCase {
	return &AdminUseCase{repo}
}

func (a *AdminUseCase) AdminSignUp(admin models.Admin) (uint, error) {
	valid := helper.IsValidEmail(admin.Email)
	if !valid {
		return 0, errors.New("please enter a valid email")
	}
	exist := a.Repo.IsAdminExist(admin.Email)
	if exist {
		return 0, errors.New("already existing email id")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return 0, errors.New("error in hashing password")
	}
	admin.Password = string(hashPassword)
	id, err := a.Repo.Save(admin)
	if err != nil {
		return 0, errors.New("error in saving admin details")
	}
	return id, nil
}

func (a *AdminUseCase) AdminLogin(admin models.Admin) (response.Token, error) {
	valid := helper.IsValidEmail(admin.Email)
	if !valid {
		return response.Token{}, errors.New("please enter a valid email")
	}
	exist := a.Repo.IsAdminExist(admin.Email)
	if !exist {
		return response.Token{}, errors.New("admin with this email id does not exist")
	}
	adminDetails, err := a.Repo.FetchAdmin(admin.Email)
	if err != nil {
		return response.Token{}, errors.New("error in fetching admin details")
	}
	err = bcrypt.CompareHashAndPassword([]byte(adminDetails.Password), []byte(admin.Password))
	if err != nil {
		return response.Token{}, errors.New("password mismatch")
	}
	accessToken,refreshToken, err := helper.GenerateAdminToken(adminDetails.Id)
	if err != nil {
		return response.Token{}, errors.New("admin token generation failed")
	}
	token:=response.Token{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
	return token, nil
}


func (a *AdminUseCase) BlockUser(id uint)(uint,error){
	block,err:=a.Repo.IsUserBlocked(id)
	if err!=nil{
		return 0,errors.New("error in connecting with database")
	}
	if block{
		return 0,errors.New("user is already blocked")
	}
	res,err:=a.Repo.BlockUser(id)
	if err!=nil{
		return 0,errors.New("error in connecting with database")
	}
	return res,nil
}

func (a *AdminUseCase) UnBlockUser(id uint)(uint,error){
	block,err:=a.Repo.IsUserBlocked(id)
	if err!=nil{
		return 0,errors.New("error in connecting with database")
	}
	if !block{
		return 0,errors.New("user is not blocked")
	}
	res,err:=a.Repo.UnblockUser(id)
	if err!=nil{
		return 0,errors.New("error in connecting with database")
	}
	return res,nil
}

func (a *AdminUseCase) GetUsers(page,count int)([]response.User,error){
	res,err:=a.Repo.GetUsers(page,count)
	if err!=nil{
		return nil,errors.New("error in fetching data ")
	}
	return res,nil
}