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