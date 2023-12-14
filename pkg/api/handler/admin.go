package handler

import (
	"net/http"

	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	UseCase useCaseInterface.AdminUseCase
}

func NewAdminHandler(usecase useCaseInterface.AdminUseCase) handlerInterface.AdminHandler {
	return &AdminHandler{usecase}
}

func (a *AdminHandler) SignUp(c *gin.Context) {
	var admin models.Admin

	if err := c.BindJSON(&admin); err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "data given is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	id,err:=a.UseCase.AdminSignUp(admin)
	if err!=nil{
		errRes:=response.MakeResponse(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusCreated,"successfully created admin",id,nil)
	c.JSON(http.StatusCreated,succRes)
}

func (a *AdminHandler) Login(c *gin.Context){
	var admin models.Admin
	if err:=c.BindJSON(&admin);err!=nil{
		errRes:=response.MakeResponse(http.StatusBadRequest,"provided data is not in required format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	token,err:=a.UseCase.AdminLogin(admin)
	if err!=nil{
		errRes:=response.MakeResponse(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusOK,"successfully login",token,nil)
	c.JSON(http.StatusOK,succRes)
}
