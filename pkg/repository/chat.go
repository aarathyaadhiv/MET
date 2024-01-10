package repository

import (
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/mongo"
)


type ChatRepository struct{
	DB *mongo.Client
}

func NewChatRepository(db *mongo.Client)interfaces.ChatRepository{
	return &ChatRepository{db}
}

func (c *ChatRepository) CreateChatRoom()(){
	
}