package handler

import (
	"fmt"
	"net/http"

	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type client struct {
	ChatId primitive.ObjectID
	UserId uint
}

type VideoCall struct {
	CallId      string
	CallerId    uint
	RecipientId uint
	Status      string
}

var (
	connection = make(map[*websocket.Conn]*client)
	user       = make(map[uint]*websocket.Conn)
	videoCall=make(map[string]*VideoCall)
)

type ChatHandler struct {
	UseCase useCaseInterface.ChatUseCase
}

func NewChatHandler(usecase useCaseInterface.ChatUseCase) handlerInterface.ChatHandler {
	return &ChatHandler{UseCase: usecase}
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

//only for without websocket testing
// // @Summary Send a message in a chat
// // @Description Sends a message in the specified chat.
// // @ID sendMessage
// // @Tags Chat
// // @Accept json
// // @Produce json
// // @Param chatId path string true "chatID" format(objectId)
// // @Param message body models.Message true "Message object"
// // @Success 200 {object} response.Response{} "Successfully sent message"
// // @Failure 400 {object} response.Response{} "Bad Request"
// // @Failure 401 {object} response.Response{} "Unauthorized"
// // @Failure 500 {object} response.Response{} "Internal Server Error"
// // @Router /chat/{chatId}/message [post]
// func (t *ChatHandler) SendMessage(c *gin.Context) {
// 	chatId, err := primitive.ObjectIDFromHex(c.Param("chatId"))
// 	if err != nil {
// 		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}
// 	id, ok := c.Get("userId")
// 	if !ok {
// 		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, "error in retrieving user id")
// 		c.JSON(http.StatusUnauthorized, errRes)
// 		return
// 	}
// 	var message models.Message
// 	if err := c.BindJSON(&message); err != nil {
// 		errRes := response.MakeResponse(http.StatusBadRequest, "data is not in required format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}
// 	res, err := t.UseCase.SaveMessage(chatId, id.(uint), message.Message)
// 	if err != nil {
// 		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errRes)
// 		return
// 	}
// 	succRes := response.MakeResponse(http.StatusOK, "successfully sent message", res, nil)
// 	c.JSON(http.StatusOK, succRes)
// }

//only for without websocket testing ends here

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

// @Summary Handle WebSocket connection for chat
// @Description Handles WebSocket connection for chat and processes incoming messages
// @ID chat-websocket
// @Tags Chat
// @Accept json
// @Produce json
// @Param chatId path string true "Chat ID"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 401 {object} response.Response{}
// @Router /ws/{chatId} [get]
func (t *ChatHandler) Chat(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "data is not in required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.MakeResponse(http.StatusUnauthorized, "unauthorised", nil, "error in retrieving user id")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}
	chatId, err := primitive.ObjectIDFromHex(c.Param("chatId"))
	if err != nil {
		errRes := response.MakeResponse(http.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	connection[conn] = &client{ChatId: chatId, UserId: id.(uint)}
	user[id.(uint)] = conn

	go func() {

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			userId := connection[conn].UserId
			chatID := connection[conn].ChatId
			_, err = t.UseCase.SaveMessage(chatID, userId, string(msg))
			if err != nil {
				fmt.Println("error in saving message")
				break
			}
			conn.WriteMessage(websocket.TextMessage, msg)
			recipient, err := t.UseCase.FetchRecipient(chatID, userId)
			if err != nil {
				fmt.Println("error in fetching recipient id")
				break
			}
			if value, ok := user[recipient]; ok {
				err = value.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					delete(connection, value)
					delete(user, recipient)
				}
			}
		}
	}()
}


func (t *ChatHandler) VideoCall(c *gin.Context) {
	peerConnectionConfig := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	peerConnection, err := webrtc.NewPeerConnection(peerConnectionConfig)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	peerConnection.OnICEConnectionStateChange(func(is webrtc.ICEConnectionState) { fmt.Printf("connection state has changed %s /n", is.String()) })

	offer := webrtc.SessionDescription{}

	peerConnection.SetRemoteDescription(offer)

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		errRes := response.MakeResponse(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	gatherComplete:=webrtc.GatheringCompletePromise(peerConnection)

	peerConnection.SetLocalDescription(answer)
	<-gatherComplete
}
