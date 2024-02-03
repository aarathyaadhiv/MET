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

type SubscriptionHandler struct {
	Usecase useCaseInterface.SubscriptionUseCase
}

func NewSubscriptionHandler(usecase useCaseInterface.SubscriptionUseCase) handlerInterface.SubscriptionHandler {
	return &SubscriptionHandler{Usecase: usecase}
}

// @Summary Add a new subscription
// @Description Add a new subscription to the system
// @Tags Subscription Management
// @Accept json
// @Produce json
// @Param subscription body models.Subscription true "Subscription object in JSON format"
// @Success 200 {object} response.Response{} "success"
// @Failure 400 {object} response.Response{} "bad request"
// @Failure 500 {object} response.Response{} "internal server error"
// @Router /admin/subscription [post]
func (s *SubscriptionHandler) Add(c *gin.Context) {
	var subscription models.Subscription
	if err := c.BindJSON(&subscription); err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "provided data is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res, err := s.Usecase.Add(subscription)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully added plan", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Update an existing subscription
// @Description Update an existing subscription in the system
// @Tags Subscription Management
// @Accept json
// @Produce json
// @Param subscriptionId path int true "Subscription ID" format(int64) minimum(1) "ID of the subscription to update"
// @Param subscription body models.Subscription true "Subscription object in JSON format"
// @Success 200 {object} response.Response{} "success"
// @Failure 400 {object} response.Response{} "bad request"
// @Failure 500 {object} response.Response{} "internal server error"
// @Router /admin/subscription/{subscriptionId} [put]
func (s *SubscriptionHandler) Update(c *gin.Context) {
	var subscription models.Subscription
	if err := c.BindJSON(&subscription); err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "provided data is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	sID, err := strconv.Atoi(c.Param("subscriptionId"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res, err := s.Usecase.Update(subscription, uint(sID))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully updated plan", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Activate or deactivate a subscription
// @Description Activate or deactivate a subscription in the system. Use 'true' to activate and 'false' to deactivate.
// @Tags Subscription Management
// @Accept json
// @Produce json
// @Param subscriptionId path int true "Subscription ID" format(int64) minimum(1) "ID of the subscription to activate or deactivate"
// @Param activate query boolean false "Activate the subscription (use 'true' for activation, 'false' for deactivation)"
// @Success 200 {object} response.Response{} "success"
// @Failure 400 {object} response.Response{} "bad request"
// @Failure 500 {object} response.Response{} "internal server error"
// @Router /admin/subscription/{subscriptionId} [patch]
func (s *SubscriptionHandler) ActivateOrDeactivate(c *gin.Context) {
	sID, err := strconv.Atoi(c.Param("subscriptionId"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	activate, err := strconv.ParseBool(c.DefaultQuery("activate", "false"))

	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "boolean conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if activate {
		res, err := s.Usecase.Activate(uint(sID))
		if err != nil {
			errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
			c.JSON(http.StatusInternalServerError, errRes)
			return
		}
		succRes := response.MakeResponse(http.StatusOK, "successfully activated plan", res, nil)
		c.JSON(http.StatusOK, succRes)
		return
	}
	res, err := s.Usecase.Dectivate(uint(sID))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully deactivated plan", res, nil)
	c.JSON(http.StatusOK, succRes)

}

// @Summary Get all subscriptions
// @Description Get a list of all subscriptions in the system
// @Tags Subscription Management
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{} "success"
// @Failure 500 {object} response.Response{} "internal server error"
// @Router /admin/subscription [get]
func (s *SubscriptionHandler) Get(c *gin.Context) {
	res, err := s.Usecase.Get()
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing plans", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Get a subscription by ID
// @Description Get details of a subscription by its ID
// @Tags Subscription Management
// @Accept json
// @Produce json
// @Param subscriptionId path int true "Subscription ID" format(int64) minimum(1) "ID of the subscription to retrieve"
// @Success 200 {object} response.Response{} "success"
// @Failure 400 {object} response.Response{} "bad request"
// @Failure 500 {object} response.Response{} "internal server error"
// @Router /admin/subscription/{subscriptionId} [get]
func (s *SubscriptionHandler) GetById(c *gin.Context) {
	sID, err := strconv.Atoi(c.Param("subscriptionId"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res, err := s.Usecase.GetById(uint(sID))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing plan", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Get subscriptions for users
// @Description Get a list of subscriptions for users in the system
// @Tags Subscription
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{} "success"
// @Failure 500 {object} response.Response{} "internal server error"
// @Router /subscription [get]
func (s *SubscriptionHandler) GetToUsers(c *gin.Context) {
	res, err := s.Usecase.GetToUsers()
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing plans", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Get a subscription for users by ID
// @Description Get details of a subscription for users by its ID
// @Tags Subscription
// @Accept json
// @Produce json
// @Param subscriptionId path int true "Subscription ID" format(int64) minimum(1) "ID of the subscription to retrieve"
// @Success 200 {object} response.Response{} "success"
// @Failure 400 {object} response.Response{} "bad request"
// @Failure 500 {object} response.Response{} "internal server error"
// @Router /subscription/{subscriptionId} [get]
func (s *SubscriptionHandler) GetByIdToUsers(c *gin.Context) {
	sID, err := strconv.Atoi(c.Param("subscriptionId"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := s.Usecase.GetByIdToUsers(uint(sID))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing plan", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Add an order for a subscription
// @Description Add an order for a subscription in the system
// @Tags Subscription
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param subscriptionId path int true "Subscription ID" format(int64) minimum(1) "ID of the subscription to order"
// @Success 200 {object} response.Response{} "success"
// @Failure 400 {object} response.Response{} "bad request"
// @Failure 401 {object} response.Response{} "unauthorized"
// @Failure 500 {object} response.Response{} "internal server error"
// @Router /subscription/order/{subscriptionId} [post]
func (s *SubscriptionHandler) AddOrder(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, "error in retrieving user id")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	sID, err := strconv.Atoi(c.Param("subscriptionId"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := s.Usecase.AddOrder(uint(sID), id.(uint))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully placed order", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Make a payment for a subscription
// @Description Make a payment for a subscription using the provided order ID
// @Tags Subscription
// @Accept json
// @Produce json
// @Param orderId path int true "Order ID"
// @Success 200 {object} response.Response{} "Payment successful"
// @Failure 400 {object} response.Response{} "Invalid input or conversion failure"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /subscription/payment/{orderId} [get]
func (s *SubscriptionHandler) MakePayment(c *gin.Context){
	orderId, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res,err:=s.Usecase.MakePayment(uint(orderId))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	c.HTML(http.StatusOK,"payment.html",res)
}

// @Summary Verify a payment for a subscription
// @Description Verify a payment for a subscription using the provided order ID, payment ID, and signature
// @Tags Subscription
// @Accept json
// @Produce json
// @Param order_id query int true "Order ID"
// @Param payment_id query string true "Payment ID"
// @Param signature query string true "Signature"
// @Success 200 {object} response.Response{} "Successfully verified payment"
// @Failure 400 {object} response.Response{} "Invalid input or conversion failure"
// @Failure 500 {object} response.Response{} "Internal server error"
// @Router /subscription/payment-success [get]
func (s *SubscriptionHandler) VerifyPayment(c *gin.Context){
	orderId,err:=strconv.Atoi(c.Query("order_id"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	paymentId:=c.Query("payment_id")
	
	signature:=c.Query("signature")
	
	err=s.Usecase.VerifyPayment(uint(orderId),signature,paymentId)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully verified payment", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Get orders for a user
// @Description Get orders associated with the authenticated user
// @Tags Subscription
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "successfully showing orders"
// @Failure 401 {object} response.Response{} "unauthorised"
// @Failure 500 {object} response.Response{} "internal server error"
// @Router /subscription/orders [get]
func (s *SubscriptionHandler) GetOrders(c *gin.Context){
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, "error in retrieving user id")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	res,err:=s.Usecase.GetOrders(id.(uint))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing orders", res, nil)
	c.JSON(http.StatusOK, succRes)
}