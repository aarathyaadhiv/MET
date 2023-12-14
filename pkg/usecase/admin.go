package usecase

import (
	"errors"

	"github.com/aarathyaadhiv/met/pkg/helper"
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
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

func (a *AdminUseCase) AdminLogin(admin models.Admin) (string, error) {
	valid := helper.IsValidEmail(admin.Email)
	if !valid {
		return "", errors.New("please enter a valid email")
	}
	exist := a.Repo.IsAdminExist(admin.Email)
	if !exist {
		return "", errors.New("admin with this email id does not exist")
	}
	adminDetails, err := a.Repo.FetchAdmin(admin.Email)
	if err != nil {
		return "", errors.New("error in fetching admin details")
	}
	err = bcrypt.CompareHashAndPassword([]byte(adminDetails.Password), []byte(admin.Password))
	if err != nil {
		return "", errors.New("password mismatch")
	}
	token, err := helper.GenerateAdminToken(adminDetails.Id)
	if err != nil {
		return "", errors.New("admin token generation failed")
	}
	return token, nil
}
