package useCaseInterface

import (
	"github.com/aarathyaadhiv/met/pkg/domain"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatUseCase interface {
	GetAllChats(userId uint) ([]response.ChatResponse, error)
	GetMessages(chatId primitive.ObjectID) ([]domain.Messages, error)
	SaveMessage(chatId primitive.ObjectID,senderId uint,message string)(primitive.ObjectID,error)
	ReadMessage(userId uint,chatId primitive.ObjectID)(int64,error)
	FetchRecipient(chatId primitive.ObjectID,userId uint)(uint,error)
}
