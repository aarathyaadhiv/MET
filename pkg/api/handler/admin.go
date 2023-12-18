package handler

import (
	"net/http"
	"strconv"

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
	id, err := a.UseCase.AdminSignUp(admin)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusCreated, "successfully created admin", id, nil)
	c.JSON(http.StatusCreated, succRes)
}

func (a *AdminHandler) Login(c *gin.Context) {
	var admin models.Admin
	if err := c.BindJSON(&admin); err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "provided data is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	token, err := a.UseCase.AdminLogin(admin)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	c.SetCookie("accessAdminToken",token.AccessToken,4500,"","",false,true)
	c.SetCookie("refreshAdminToken",token.RefreshToken,4500,"","",false,true)
	succRes := response.MakeResponse(http.StatusOK, "successfully login", token, nil)
	c.JSON(http.StatusOK, succRes)
}

func (a *AdminHandler) BlockOrUnBlock(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "boolean conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	block, err := strconv.ParseBool(c.DefaultQuery("block", "false"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "boolean conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if block {
		res, err := a.UseCase.BlockUser(uint(id))
		if err != nil {
			errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
			c.JSON(http.StatusInternalServerError, errRes)
			return
		}
		succRes := response.MakeResponse(http.StatusOK, "successfully blocked user", res, nil)
		c.JSON(http.StatusOK, succRes)
		return
	}
	res, err := a.UseCase.UnBlockUser(uint(id))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully unblocked user", res, nil)
	c.JSON(http.StatusOK, succRes)
}

func (a *AdminHandler) GetUsers(c *gin.Context){
	res,err:=a.UseCase.GetUsers()
	if err!=nil{
		errRes:=response.MakeResponse(http.StatusInternalServerError,"internal serverr error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusOK,"successfully showing users",res,nil)
	c.JSON(http.StatusOK,succRes)
}
