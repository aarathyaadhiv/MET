package handler

import (
	"net/http"

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