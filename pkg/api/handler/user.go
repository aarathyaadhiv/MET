package handler

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"
	"strings"
	"time"

	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UseCase useCaseInterface.UserUseCase
}

func NewUserHandler(useCase useCaseInterface.UserUseCase) handlerInterface.UserHandler {
	return &UserHandler{useCase}
}

// @Summary User Login
// @Description sending otp to the given phone number
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param  sendOtp body models.OtpRequest true "sendOtp"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /sendOtp [post]
func (u *UserHandler) SendOtp(c *gin.Context) {
	var sendOtp models.OtpRequest
	if err := c.BindJSON(&sendOtp); err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "data provided is not in the correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res,err := u.UseCase.SendOtp(sendOtp.PhNo)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "error in sending otp", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully sent otp", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Verify OTP
// @Description Verify OTP for user authentication and generate access and refresh tokens
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param request body models.OtpVerify true "OTP verification details"
// @Success 200 {object} response.Response{} "Successfully verified existing user"
// @Success 201 {object} response.Response{} "Successfully created user"
// @Failure 400 {object} response.Response{} "Data is not in required format"
// @Failure 500 {object} response.Response{} "Error in verifying OTP"
// @Router /verify [post]
func (u *UserHandler) VerifyOtp(c *gin.Context) {
	var verify models.OtpVerify
	if err := c.BindJSON(&verify); err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "data is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	exist, token, err := u.UseCase.VerifyOtp(verify)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "error in verifying otp", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	c.SetCookie("accessToken", token.AccessToken, 172800, "", "", false, true)
	c.SetCookie("refreshToken", token.RefreshToken, 172800, "", "", false, true)
	if exist {
		succRes := response.MakeResponse(http.StatusOK, "successfully verified existing user", token, nil)
		c.JSON(http.StatusOK, succRes)
		return
	}
	succRes := response.MakeResponse(http.StatusCreated, "successfully created user", token, nil)
	c.JSON(http.StatusCreated, succRes)
}

// @Summary Add user profile details
// @Description Add user profile details including name, date of birth, gender, location, bio, interests, and image
// @Tags User Profile
// @Security ApiKeyAuth
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Name"
// @Param dob formData string true "Date of Birth (YYYY-MM-DD)"
// @Param genderId formData integer true "Gender ID"
// @Param city formData string true "City"
// @Param country formData string true "Country"
// @Param longitude formData float64 true "Longitude"
// @Param lattitude formData float64 true "Latitude"
// @Param bio formData string true "Bio"
// @Param interests formData string true "Interests (comma-separated IDs)"
// @Param images[] formData file true "Image"
// @Success 200 {object} response.Response{} "Successfully added user details"
// @Failure 400 {object} response.Response{} "Data is not in the required format"
// @Failure 401 {string} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /profile [post]
func (u *UserHandler) AddProfile(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthourized", nil, errors.New("error in getting id"))
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	var profile models.Profile
	profile.Name = c.Request.FormValue("name")

	parsedDob, err := time.Parse("2006-01-02", c.Request.FormValue("dob"))
	if err != nil {

		errRes := response.MakeResponse(http.StatusBadRequest, "dob is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	profile.Dob = parsedDob

	genderId, err := strconv.Atoi(c.Request.FormValue("genderId"))
	if err != nil {

		errRes := response.MakeResponse(http.StatusBadRequest, "gender is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	profile.GenderId = uint(genderId)

	profile.City = c.Request.FormValue("city")

	profile.Country = c.Request.FormValue("country")

	profile.Longitude, err = strconv.ParseFloat(c.Request.FormValue("longitude"), 64)
	if err != nil {

		errRes := response.MakeResponse(http.StatusBadRequest, "longitude is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	profile.Lattitude, err = strconv.ParseFloat(c.Request.FormValue("lattitude"), 64)
	if err != nil {

		errRes := response.MakeResponse(http.StatusBadRequest, "lattitude is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	profile.Bio = c.Request.FormValue("bio")

	interest := c.Request.FormValue("interests")
	var interests []uint
	value := strings.Split(interest, ",")
	for _, v := range value {
		val, err := strconv.Atoi(v)
		if err != nil {

			errRes := response.MakeResponse(http.StatusBadRequest, "interest is not in required format", nil, err.Error())
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		interests = append(interests, uint(val))
	}

	profile.Interests = interests

	image, err := c.MultipartForm()

	if err != nil {

		errRes := response.MakeResponse(http.StatusBadRequest, "data is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	profile.Image = image

	res, err := u.UseCase.AddProfile(profile, id.(uint))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully added user details", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Get user profile details
// @Description Retrieve user profile details based on the user ID
// @Tags User Profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Successfully showing profile"
// @Failure 401 {object} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /profile [get]
func (u *UserHandler) GetProfile(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, errors.New("error in retrieving user id"))
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	profile, err := u.UseCase.ShowProfile(id.(uint))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing profile", profile, nil)
	c.JSON(http.StatusOK, succRes)

}

// @Summary Update user profile
// @Description Update user profile information including name,phone number, city, country, bio, interests
// @Tags User Profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body models.UpdateUser true "user object in JSON format"
// @Success 200 {object} response.Response{} "Successfully updated user profile"
// @Failure 400 {object} response.Response{} "Bad request or invalid data format"
// @Failure 401 {object} response.Response{} "Unauthorized access"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /profile [put]
func (u *UserHandler) UpdateUser(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, errors.New("error in retrieving user id"))
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	var user models.UpdateUser

	if err := c.BindJSON(&user); err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "provided data is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := u.UseCase.UpdateUser(user, id.(uint))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully updated profile", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Update user Image
// @Description Update user profile image
// @Tags User Profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param images formData file false "Images to upload"
// @Success 200 {object} response.Response{} "Successfully updated user profile image"
// @Failure 400 {object} response.Response{} "Bad request or invalid data format"
// @Failure 401 {object} response.Response{} "Unauthorized access"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /profile [patch]
func (u *UserHandler) UpdateImage(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, errors.New("error in retrieving user id"))
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	image, err := c.MultipartForm()
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "data is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res, err := u.UseCase.UpdateImage(id.(uint), image)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully updated profile image", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Update user preferences
// @Description Update user preferences such as distance and other criteria
// @Tags User Preferences
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param preference body models.Preference true "User preference details"
// @Success 200 {object} response.Response{} "Successfully updated preference"
// @Failure 400 {object} response.Response{} "Data is not in the required format"
// @Failure 401 {string} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /preference [put]
func (u *UserHandler) UpdatePreference(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, "error in retrieving user id")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	var preference models.Preference
	fmt.Println("preference", preference)
	if err := c.BindJSON(&preference); err != nil {
		fmt.Println("preference2", preference)
		fmt.Println("err", err)
		errRes := response.MakeResponse(http.StatusBadRequest, "data is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res, err := u.UseCase.UpdatePreference(id.(uint), preference)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully updated preference", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Get user preferences
// @Description Retrieve user preferences such as distance and other criteria
// @Tags User Preferences
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{} "Successfully retrieved preference"
// @Failure 401 {string} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /preference [get]
func (u *UserHandler) GetPreference(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, "error in retrieving user id")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	res, err := u.UseCase.GetPreference(id.(uint))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing preference", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Show  interests
// @Description Returns interests of the user if value of user is true, otherwise gives all interests
// @Tags Profile
// @Security ApiKeyAuth
// @Produce json
// @Param user query boolean false "filter for particular user (default: false)"
// @Success 200 {object} response.Response{}
// @Failure 401 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /interests [get]
func (u *UserHandler) Interests(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthourized", nil, "error in getting id")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	user, err := strconv.ParseBool(c.DefaultQuery("user", "false"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res, err := u.UseCase.Interests(id.(uint), user)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing interests ", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Show genders
// @Description Returns available genders list
// @Tags Profile
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 401 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /genders [get]
func (u *UserHandler) Genders(c *gin.Context) {
	res, err := u.UseCase.Gender()
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing genders", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Logout user
// @Description Remove access and refresh tokens from cookies
// @Tags User Profile
// @Produce json
// @Router /logout [post]
func (u *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "", "", false, true)
	c.SetCookie("refreshToken", "", -1, "", "", false, true)
}

// @Summary Delete user account
// @Description Deletes the user account associated with the authenticated user
// @Tags User Profile
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{}
// @Failure 401 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /profile [delete]
func (u *UserHandler) DeleteAccount(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, "error in retrieving user id")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	res, err := u.UseCase.DeleteAccount(id.(uint))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	u.Logout(c)
	succRes := response.MakeResponse(http.StatusOK, "successfully deleted account", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Send OTP to Update PhNo 
// @Description sending otp to the given phone number to replace the old number with the given number
// @Tags User Profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param  sendOtp body models.OtpRequest true "sendOtp"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /profile/phNo/sendOtp [post]
func (u *UserHandler) UpdatePhNoSendOTP(c *gin.Context) {
	var sendOtp models.OtpRequest
	if err := c.BindJSON(&sendOtp); err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "data provided is not in the correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res, err := u.UseCase.SendOtp(sendOtp.PhNo)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "error in sending otp", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully sent otp", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Verify OTP to Update PhNo
// @Description Verify OTP for updating user saved phNo to given phNo
// @Tags User Profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body models.OtpVerify true "OTP verification details"
// @Success 200 {object} response.Response{} 
// @Failure 401 {object} response.Response{}
// @Failure 400 {object} response.Response{} 
// @Failure 500 {object} response.Response{} 
// @Router /profile/phNo/verify [patch]
func (u *UserHandler) VerifyOTPtoUpdatePhNo(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, "error in retrieving user id")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	var verify models.OtpVerify
	if err := c.BindJSON(&verify); err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "data is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res, err := u.UseCase.VerifyOTPtoUpdatePhNo(verify, id.(uint))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "error in verifying otp", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	succRes := response.MakeResponse(http.StatusOK, "successfully updated PhNo", res, nil)
	c.JSON(http.StatusCreated, succRes)
}
