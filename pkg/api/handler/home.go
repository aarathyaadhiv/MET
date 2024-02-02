package handler

import (
	"net/http"
	"strconv"

	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type HomeHandler struct {
	UseCase useCaseInterface.HomeUseCase
}

func NewHomeHandler(usecase useCaseInterface.HomeUseCase) handlerInterface.HomeHandler {
	return &HomeHandler{usecase}
}

// @Summary Display the home page for a user
// @Description Fetches the user's  home page to show the other user's profile to make match
// @Tags Home
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param page query integer false "Page number (default: 1)"
// @Param count query integer false "Number of items per page (default: 3)"
// @Param interest query boolean false "interest filter (default: false)" 
// @Param interestId query integer false "interestId" 
// @Success 200 {object} response.Response{} "Successfully fetched home page data"
// @Failure 401 {object} response.Response{} "Unauthorized: Invalid or missing authentication token"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /home [get]
func (h *HomeHandler) Home(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthourized", nil, "error in getting id")
		c.JSON(http.StatusUnauthorized, errRes)
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
	interest,err:=strconv.ParseBool(c.DefaultQuery("interest", "false"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	interestId:=c.Query("interestId")
	res, err := h.UseCase.HomePage(id.(uint), page, count,interest,interestId)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing home page", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Show user interests
// @Description Returns interests of the user
// @Tags Home
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 401 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /interests [get]
func (h *HomeHandler) ShowInterests(c *gin.Context){
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthourized", nil, "error in getting id")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	res,err:=h.UseCase.Interests(id.(uint))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing interests of the user", res, nil)
	c.JSON(http.StatusOK, succRes)
}
