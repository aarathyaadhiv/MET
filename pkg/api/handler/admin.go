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

// @Summary Create a new admin
// @Description Create a new admin with provided details
// @Tags Admin Authentication
// @Accept json
// @Produce json
// @Param request body models.Admin true "Admin details"
// @Success 201 {object} response.Response{} "Successfully created admin"
// @Failure 400 {object} response.Response{} "Data given is not in required format"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /admin/signUp [post]
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
// @Summary Log in as an admin
// @Description Log in as an admin with provided credentials
// @Tags Admin Authentication
// @Accept json
// @Produce json
// @Param request body models.Admin true "Admin login credentials"
// @Success 200 {object} response.Response{} "Successfully logged in"
// @Failure 400 {object} response.Response{} "Provided data is not in required format"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /admin/login [post]
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
// @Summary Block or unblock a user
// @Description Block or unblock a user based on the provided ID and block status
// @Tags User Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path integer true "User ID"
// @Param block query boolean true "Block status: true to block, false to unblock"
// @Success 200 {object} response.Response{} "Successfully blocked/unblocked user"
// @Failure 400 {object} response.Response{} "Boolean conversion failed"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /admin/users/{id} [patch]
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
// @Summary Get all users to admin
// @Description Retrieve all users
// @Tags User Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query integer false "Page number (default: 1)"
// @Param count query integer false "Number of items per page (default: 3)"
// @Success 200 {object} response.Response{} "Successfully retrieved users"
// @Failure 400 {object} response.Response{} "int conversion failed"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /admin/users [get]
func (a *AdminHandler) GetUsers(c *gin.Context){
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	count, err := strconv.Atoi(c.DefaultQuery("count", "3"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res,err:=a.UseCase.GetUsers(page,count)
	if err!=nil{
		errRes:=response.MakeResponse(http.StatusInternalServerError,"internal serverr error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusOK,"successfully showing users",res,nil)
	c.JSON(http.StatusOK,succRes)
}
