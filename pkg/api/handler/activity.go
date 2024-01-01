package handler

import (
	"errors"
	"net/http"
	"strconv"

	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"github.com/gin-gonic/gin"
)


type ActivityHandler struct{
	Usecase useCaseInterface.ActivityUseCase
}

func NewActivityHandler(activity useCaseInterface.ActivityUseCase)handlerInterface.ActivityHandler{
	return &ActivityHandler{activity}
}
// @Summary Like an user
// @Description Likes an user based on the provided ID
// @Tags Likes
// @Accept json
// @Produce json
// @Param id path int true "Item ID to like"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Successfully liked"
// @Failure 400 {object} response.Response{} "String conversion failed"
// @Failure 401 {object} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /like/{id} [post]
func (a *ActivityHandler) Like(c *gin.Context){
	userId,ok:=c.Get("userId")
	if !ok{
		errRes:=response.MakeResponse(http.StatusUnauthorized,"unauthourized",nil,errors.New("error in getting id"))
		c.JSON(http.StatusUnauthorized,errRes)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res,err:=a.Usecase.Like(uint(id),userId.(uint))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusOK,"successfully liked",res,nil)
	c.JSON(http.StatusOK,succRes)
}
// @Summary Get liked users for a user
// @Description Retrieves liked users for the authenticated user
// @Tags Likes
// @Accept json
// @Produce json
// @Param page query int false "Page number (default is 1)"
// @Param count query int false "Number of items per page (default is 3)"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Successfully retrieved liked items"
// @Failure 400 {object} response.Response{} "Bad request"
// @Failure 401 {object} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /like [get]
func (a *ActivityHandler) GetLike(c *gin.Context){
	id,ok:=c.Get("userId")
	if !ok{
		errRes:=response.MakeResponse(http.StatusUnauthorized,"unauthorised",nil,errors.New("error in retrieving user id"))
		c.JSON(http.StatusUnauthorized,errRes)
		return
	}
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
	res,err:=a.Usecase.GetLike(page,count,id.(uint))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusOK,"successfully showing likes",res,nil)
	c.JSON(http.StatusOK,succRes)
}
// @Summary Unmatch with a user
// @Description Unmatches with a user based on the provided ID
// @Tags Match
// @Accept json
// @Produce json
// @Param id path int true "User ID to unmatch"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Successfully unmatched"
// @Failure 400 {object} response.Response{} "String conversion failed"
// @Failure 401 {object} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /unmatch/{id} [delete]
func (a *ActivityHandler) Unmatch(c *gin.Context){
	userId,ok:=c.Get("userId")
	if !ok{
		errRes:=response.MakeResponse(http.StatusUnauthorized,"unauthourized",nil,errors.New("error in getting id"))
		c.JSON(http.StatusUnauthorized,errRes)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res,err:=a.Usecase.UnMatch(userId.(uint),uint(id))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusOK,"successfully unmatched",res,nil)
	c.JSON(http.StatusOK,succRes)
}

// @Summary Get matched items for a user
// @Description Retrieves matched items for the authenticated user
// @Tags Match
// @Accept json
// @Produce json
// @Param page query int false "Page number (default is 1)"
// @Param count query int false "Number of items per page (default is 3)"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Successfully retrieved matched items"
// @Failure 400 {object} response.Response{} "Bad request"
// @Failure 401 {object} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /match [get]
func (a *ActivityHandler)GetMatch(c *gin.Context){
	id,ok:=c.Get("userId")
	if !ok{
		errRes:=response.MakeResponse(http.StatusUnauthorized,"unauthorised",nil,errors.New("error in retrieving user id"))
		c.JSON(http.StatusUnauthorized,errRes)
		return
	}
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
	res,err:=a.Usecase.GetMatch(id.(uint),page,count)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusOK,"successfully showing matches",res,nil)
	c.JSON(http.StatusOK,succRes)
}
// @Summary Report a user
// @Description Reports a user based on the provided ID
// @Tags Activity
// @Accept json
// @Produce json
// @Param id path int true "User ID to report"
// @Param report body models.Report true "Report details"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Successfully reported"
// @Failure 400 {object} response.Response{} "String conversion failed" or "Data provided is not in the correct format"
// @Failure 401 {object} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /report/{id} [post]
func (a *ActivityHandler) Report(c *gin.Context){
	userId,ok:=c.Get("userId")
	if !ok{
		errRes:=response.MakeResponse(http.StatusUnauthorized,"unauthourized",nil,errors.New("error in getting id"))
		c.JSON(http.StatusUnauthorized,errRes)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var report models.Report
	if err:=c.BindJSON(&report);err!=nil{
		errRes:=response.MakeResponse(http.StatusBadRequest,"data provided is not in the correct format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	res,err:=a.Usecase.Report(uint(id),userId.(uint),report.Message)
	
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusOK,"successfully reported",res,nil)
	c.JSON(http.StatusOK,succRes)
}
// @Summary Block a user
// @Description Blocks a user based on the provided ID
// @Tags Activity
// @Accept json
// @Produce json
// @Param id path int true "User ID to block"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "Successfully blocked user"
// @Failure 400 {object} response.Response{} "String conversion failed"
// @Failure 401 {object} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /block/{id} [post]
func (a *ActivityHandler) BlockUser(c *gin.Context){
	userId,ok:=c.Get("userId")
	if !ok{
		errRes:=response.MakeResponse(http.StatusUnauthorized,"unauthourized",nil,errors.New("error in getting id"))
		c.JSON(http.StatusUnauthorized,errRes)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res,err:=a.Usecase.BlockUser(userId.(uint),uint(id))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes:=response.MakeResponse(http.StatusOK,"successfully blocked user",res,nil)
	c.JSON(http.StatusOK,succRes)
}