package handler

import (
	"net/http"

	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatHandler struct {
	UseCase useCaseInterface.ChatUseCase
}

func NewChatHandler(usecase useCaseInterface.ChatUseCase) handlerInterface.ChatHandler {
	return &ChatHandler{usecase}
}

// @Summary Get user's chats
// @Description Get all chats associated with the authenticated user.
// @Tags Chat
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{} "successfully showing chats of the user"
// @Failure 401 {object} response.Response{} "unauthorised"
// @Failure 500 {object} response.Response{} "internal server error"
// @Router /chat [get]
func (t *ChatHandler) GetChats(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, "error in retrieving user id")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	res, err := t.UseCase.GetAllChats(id.(uint))
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing chats of the user", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Get messages in a chat
// @Description Get all messages in a chat based on the provided chatId.
// @Tags Chat
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param chatId path string true "chatID" format(objectId)
// @Success 200 {object} response.Response{} "successfully showing messages in the given chatId"
// @Failure 400 {object} response.Response{} "string converrsion failed"
// @Failure 500 {object} response.Response{} "internal server error"
// @Router /chat/{chatId}/message [get]
func (t *ChatHandler) GetMessages(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("chatId"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res, err := t.UseCase.GetMessages(id)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully showing messages in the given chatId", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Send a message in a chat
// @Description Sends a message in the specified chat.
// @ID sendMessage
// @Tags Chat
// @Accept json
// @Produce json
// @Param chatId path string true "chatID" format(objectId)
// @Param message body models.Message true "Message object"
// @Success 200 {object} response.Response{} "Successfully sent message"
// @Failure 400 {object} response.Response{} "Bad Request"
// @Failure 401 {object} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal Server Error"
// @Router /chat/{chatId}/message [post]
func (t *ChatHandler) SendMessage(c *gin.Context) {
	chatId, err := primitive.ObjectIDFromHex(c.Param("chatId"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, "error in retrieving user id")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	var message models.Message
	if err := c.BindJSON(&message); err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "data is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res, err := t.UseCase.SaveMessage(chatId, id.(uint), message.Message)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully sent message", res, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Mark messages as read
// @Description Marks specified messages as read by the user.
// @ID makeMessageRead
// @Tags Chat
// @Accept json
// @Produce json
// @Param messageIds body models.MakeReadReq true "Message IDs to mark as read"
// @Success 200 {object} response.Response{} "Successfully marked messages as read"
// @Failure 400 {object} response.Response{} "Bad Request"
// @Failure 401 {object} response.Response{} "Unauthorized"
// @Failure 500 {object} response.Response{} "Internal Server Error"
// @Router /chat/message/read [post]
func (t *ChatHandler) MakeMessageRead(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, "error in retrieving user id")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	var messageId models.MakeReadReq
	if err := c.BindJSON(&messageId); err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "data is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res, err := t.UseCase.ReadMessage(id.(uint), messageId.MessageIds)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.MakeResponse(http.StatusOK, "successfully make these messages to read", res, nil)
	c.JSON(http.StatusOK, succRes)
}
