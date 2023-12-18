package handler

import (
	"errors"
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
func (u *UserHandler) SendOtp(c *gin.Context){
	var sendOtp models.OtpRequest
	if err:=c.BindJSON(&sendOtp);err!=nil{
		errRes:=response.MakeResponse(http.StatusBadRequest,"data provided is not in the correct format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	err:=u.UseCase.SendOtp(sendOtp.PhNo)
	if err!=nil{
		errRes:=response.MakeResponse(http.StatusInternalServerError,"error in sending otp",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusOK,"successfully sent otp",nil,nil)
	c.JSON(http.StatusOK,succRes)
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
func (u *UserHandler) VerifyOtp(c *gin.Context){
	var verify models.OtpVerify
	if err:=c.BindJSON(&verify);err!=nil{
		errRes:=response.MakeResponse(http.StatusBadRequest,"data is not in required format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	exist,token,err:=u.UseCase.VerifyOtp(verify)
	if err!=nil{
		errRes:=response.MakeResponse(http.StatusInternalServerError,"error in verifying otp",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	c.SetCookie("accessToken",token.AccessToken,4500,"","",false,true)
	c.SetCookie("refreshToken",token.RefreshToken,4500,"","",false,true)
	if exist{
		succRes:=response.MakeResponse(http.StatusOK,"successfully verified existing user",token,nil)
		c.JSON(http.StatusOK,succRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusCreated,"successfully created user",token,nil)
	c.JSON(http.StatusCreated,succRes)
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
// @Param longitude formData string true "Longitude"
// @Param lattitude formData string true "Latitude"
// @Param bio formData string true "Bio"
// @Param interests formData string true "Interests (comma-separated IDs)"
// @Param image formData file true "Image"
// @Success 200 {object} response.Response{} "Successfully added user details"
// @Failure 400 {object} response.Response{} "Data is not in the required format"
// @Failure 401 {string} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /profile [post]
func (u *UserHandler)AddProfile(c *gin.Context){
	id,ok:=c.Get("userId")
	if !ok{
		errRes:=response.MakeResponse(http.StatusUnauthorized,"unauthourized",nil,errors.New("error in getting id"))
		c.JSON(http.StatusUnauthorized,errRes)
		return
	}
	var profile models.Profile
	profile.Name=c.Request.FormValue("name")
	parsedDob,err:=time.Parse("2006-01-02",c.Request.FormValue("dob"))
	if err!=nil{
		errRes:=response.MakeResponse(http.StatusBadRequest,"data is not in required format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	profile.Dob=parsedDob
	genderId,err:=strconv.Atoi(c.Request.FormValue("genderId"))
	if err!=nil{
		errRes:=response.MakeResponse(http.StatusBadRequest,"data is not in required format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	profile.GenderId=uint(genderId)
	profile.City=c.Request.FormValue("city")
	profile.Country=c.Request.FormValue("country")
	profile.Longitude=c.Request.FormValue("longitude")
	profile.Lattitude=c.Request.FormValue("lattitude")
	profile.Bio=c.Request.FormValue("bio")
	interest:=c.Request.FormValue("interests")
	var interests []uint
	value:=strings.Split(interest,",")
	for _,v:=range value{
		val,err:=strconv.Atoi(v)
		if err!=nil{
		errRes:=response.MakeResponse(http.StatusBadRequest,"data is not in required format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
		}
		interests=append(interests, uint(val))
	}

	profile.Interests=interests
	image,formErr:=c.MultipartForm()
	if formErr!=nil{
		errRes:=response.MakeResponse(http.StatusBadRequest,"data is not in required format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	profile.Image=image
	res,err:=u.UseCase.AddProfile(profile,id.(uint))
	if err!=nil{
		errRes:=response.MakeResponse(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusOK,"successfully added user details",res,nil)
	c.JSON(http.StatusOK,succRes)
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
func (u *UserHandler) GetProfile(c *gin.Context){
	id,ok:=c.Get("userId")
	if !ok{
		errRes:=response.MakeResponse(http.StatusUnauthorized,"unauthorised",nil,errors.New("error in retrieving user id"))
		c.JSON(http.StatusUnauthorized,errRes)
		return
	}
	profile,err:=u.UseCase.ShowProfile(id.(uint))
	if err!=nil{
		errRes:=response.MakeResponse(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusOK,"successfully showing profile",profile,nil)
	c.JSON(http.StatusOK,succRes)

}