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
	"github.com/aarathyaadhiv/met/pkg/utils/response"
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
	if !regex.MatchString(phNo) {
		return errors.New("the given phone number is not in the correct format")
	}
	exist, err := u.Repo.IsUserExist(phNo)
	if err != nil {
		return err
	}
	if exist {
		block, err := u.Repo.IsUserBlocked(phNo)
		if err != nil {
			return err
		}
		if block {
			return errors.New("this number is blocked for this app")
		}
	}
	phone := "+91" + phNo
	helper.TwillioSetup(u.Config.TwilioAccountSID, u.Config.TwilioAuthToken)
	_, err = helper.SendOtp(phone, u.Config.TwilioServicesId)
	if err != nil {
		fmt.Println("error here", err)
		return errors.New("error in sending otp")
	}
	return nil
}

func (u *UserUseCase) VerifyOtp(otp models.OtpVerify) (bool, response.Token, error) {
	regex := regexp.MustCompile(`^\d{10}$`)
	if !regex.MatchString(otp.PhNo) {
		return false, response.Token{}, errors.New("the given phone number is not in the correct format")
	}
	helper.TwillioSetup(u.Config.TwilioAccountSID, u.Config.TwilioAuthToken)
	err := helper.ValidateOtp(otp, u.Config.TwilioServicesId)
	if err != nil {
		return false, response.Token{}, errors.New("otp verification failed")
	}
	//user search
	exist, err := u.Repo.IsUserExist(otp.PhNo)
	if err != nil {
		return false, response.Token{}, err
	}
	var id uint
	if exist {
		id, err = u.Repo.FindByPhone(otp.PhNo)
		if err != nil {
			return false, response.Token{}, err
		}
	} else {
		id, err = u.Repo.CreateUserId(otp.PhNo)
		if err != nil {
			return false, response.Token{}, err
		}
	}

	//generate token
	accessToken, refreshToken, err := helper.GenerateUserToken(id)
	if err != nil {
		return false, response.Token{}, err
	}
	token := response.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return exist, token, nil
}

func (u *UserUseCase) AddProfile(profile models.Profile, id uint) (uint, error) {
	age := helper.CalculateAge(profile.Dob)
	user := models.ProfileSave{
		Name:      profile.Name,
		Dob:       profile.Dob,
		Age:       age,
		GenderId:  profile.GenderId,
		City:      profile.City,
		Country:   profile.Country,
		Longitude: profile.Longitude,
		Lattitude: profile.Lattitude,
		Bio:       profile.Bio,
	}
	userId, err := u.Repo.UpdateUser(id, user)
	if err != nil {
		return 0, errors.New("error in saving details")
	}
	for _, form := range profile.Image.File {
		for _, file := range form {
			url, err := helper.AddImageToS3(file)
			if err != nil {

				return 0, errors.New("error in adding image")
			}
			exist, err := u.Repo.IsImageExist(userId, url)
			if err != nil {
				return 0, errors.New("error in saving image")
			}
			if !exist {
				err = u.Repo.AddImage(userId, url)
				if err != nil {
					return 0, errors.New("error in saving image")
				}
			}

		}
	}

	for _, interest := range profile.Interests {

		err = u.Repo.AddInterest(userId, interest)
		if err != nil {
			return 0, errors.New("error in saving interest")
		}

	}
	//add preference
	minAge, maxAge := helper.MinAndMaxAge(age)
	gender := helper.Gender(profile.GenderId)
	preference := models.Preference{
		MinAge:      minAge,
		MaxAge:      maxAge,
		Gender:      gender,
		MaxDistance: 15,
	}
	err = u.Repo.AddPreference(userId, preference)
	if err != nil {
		return 0, errors.New("error in add preference")
	}
	return userId, nil
}

func (u *UserUseCase) ShowProfile(id uint) (response.Profile, error) {
	userDetails, err := u.Repo.ShowProfile(id)
	if err != nil {
		return response.Profile{}, errors.New("error in fetching user details")
	}
	images, err := u.Repo.FetchImages(id)
	if err != nil {
		return response.Profile{}, errors.New("error in fetching user images")
	}
	interests, err := u.Repo.FetchInterests(id)
	if err != nil {
		return response.Profile{}, errors.New("error in fetching user interests")
	}
	return response.Profile{
		UserDetails: userDetails,
		Image:       images,
		Interests:   interests,
	}, nil
}

func (u *UserUseCase) UpdateUser(user models.UpdateUser, id uint) (response.Id, error) {
	users := models.UpdateUserDetails{
		Name:    user.Name,
		PhNo:    user.PhNo,
		City:    user.City,
		Country: user.Country,
		Bio:     user.Bio,
	}
	err := u.Repo.UpdateUserDetails(id, users)
	if err != nil {
		return response.Id{}, errors.New("error in fetching data")
	}
	err = u.Repo.DeleteImage(id)
	if err != nil {
		return response.Id{}, errors.New("error in fetching data")
	}
	for _, form := range user.Image.File {
		for _, file := range form {
			url, err := helper.AddImageToS3(file)
			if err != nil {
				fmt.Println("err", err)
				return response.Id{}, errors.New("error in adding image")
			}
			err = u.Repo.AddImage(id, url)
			if err != nil {
				return response.Id{}, errors.New("error in saving image")
			}
		}
	}
	err = u.Repo.DeleteInterest(id)
	if err != nil {
		return response.Id{}, errors.New("error in fetching data")
	}
	for _, interest := range user.Interests {
		err = u.Repo.AddInterest(id, interest)
		if err != nil {
			return response.Id{}, errors.New("error in saving interest")
		}
	}
	return response.Id{
		Id: id,
	}, nil
}

func (u *UserUseCase) UpdatePreference(id uint, preference models.Preference) (response.Id, error) {
	userId, err := u.Repo.UpdatePreference(id, preference)
	if err != nil {
		return response.Id{}, err
	}
	return response.Id{
		Id: userId,
	}, nil
}

func (u *UserUseCase) GetPreference(id uint) (models.Preference, error) {
	res, err := u.Repo.GetPreference(id)
	if err != nil {
		return models.Preference{}, err
	}
	return res, nil
}
