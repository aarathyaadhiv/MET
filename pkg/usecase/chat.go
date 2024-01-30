package usecase

import (
	"time"

	"github.com/aarathyaadhiv/met/pkg/domain"
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatUseCase struct {
	Repo interfaces.ChatRepository
	User interfaces.UserRepository
}

func NewChatUseCase(repo interfaces.ChatRepository, user interfaces.UserRepository) useCaseInterface.ChatUseCase {
	return &ChatUseCase{Repo: repo, User: user}
}

func (c *ChatUseCase) GetAllChats(userId uint) ([]response.ChatResponse, error) {
	chats, err := c.Repo.GetAllChats(userId)
	if err != nil {
		return nil, err
	}
	chatResponses := make([]response.ChatResponse, 0)
	for _, chat := range chats {
		user, err := c.User.FetchShortDetail(chat.Users[0])
		if err != nil {
			return nil, err
		}
		chatResponse := response.ChatResponse{
			Chat: chat,
			User: user,
		}
		chatResponses = append(chatResponses, chatResponse)
	}
	return chatResponses, nil
}

func (c *ChatUseCase) GetMessages(chatId primitive.ObjectID) ([]domain.Messages, error) {
	messages, err := c.Repo.GetMessages(chatId)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (c *ChatUseCase) SaveMessage(chatId primitive.ObjectID, senderId uint, message string) (primitive.ObjectID, error) {
	messages := domain.Messages{
		SenderID:       senderId,
		ChatID:         chatId,
		MessageContent: message,
		Timestamp:      time.Now(),
	}
	res, err := c.Repo.SaveMessage(messages)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return res, nil
}

func (c *ChatUseCase) ReadMessage(userId uint, chatId primitive.ObjectID) (int64, error) {

	senderId, err := c.Repo.FetchRecipient(chatId, userId)

	if err != nil {
		return 0, err
	}

	res, err := c.Repo.ReadMessage(chatId, senderId)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (c *ChatUseCase) FetchRecipient(chatId primitive.ObjectID, userId uint) (uint, error) {
	res, err := c.Repo.FetchRecipient(chatId, userId)
	if err != nil {
		return 0, err
	}
	return res, nil
}
