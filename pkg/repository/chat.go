package repository

import (
	"context"

	"time"

	"github.com/aarathyaadhiv/met/pkg/domain"
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatRepository struct {
	ChatCollection    *mongo.Collection
	MessageCollection *mongo.Collection
}

func NewChatRepository(db *mongo.Database) interfaces.ChatRepository {
	return &ChatRepository{ChatCollection: db.Collection("chats"), MessageCollection: db.Collection("messages")}
}

func (c *ChatRepository) CreateChatRoom(user1, user2 uint) error {
	newChat := domain.Chats{
		Users:     []uint{user1, user2},
		CreatedAt: time.Now(),
	}
	_, err := c.ChatCollection.InsertOne(context.TODO(), &newChat)
	return err
}

func (c *ChatRepository) IsChatExist(user1, user2 uint) (bool, error) {
	filter := bson.M{
		"users": bson.M{"$all": []uint{user1, user2}},
	}

	var chat domain.Chats
	err := c.ChatCollection.FindOne(context.TODO(), filter).Decode(&chat)

	if err == mongo.ErrNoDocuments {

		return false, nil
	} else if err != nil {

		return false, err
	}

	return true, nil
}

func (c *ChatRepository) GetAllChats(id uint) ([]models.Chat, error) {
	// Define the filter and projection
	filter := bson.M{"users": bson.M{"$in": []uint{id}}}
	projection := bson.M{"_id": 1, "users": bson.M{"$elemMatch": bson.M{"$ne": id}}, "created_at": 1}

	// Execute the find query
	cursor, err := c.ChatCollection.Find(context.TODO(), filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var chats []models.Chat
	if err := cursor.All(context.TODO(), &chats); err != nil {
		return nil, err
	}

	return chats, nil
}

func (c *ChatRepository) GetMessages(id primitive.ObjectID) ([]domain.Messages, error) {
	ctx := context.TODO()
	filter := bson.M{"chat_id": id}
	cursor, err := c.MessageCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var messages []domain.Messages
	err = cursor.All(ctx, &messages)
	if err != nil {
		return nil, err
	}
	return messages, nil

}

func (c *ChatRepository) SaveMessage(message domain.Messages) (primitive.ObjectID, error) {
	id, err := c.MessageCollection.InsertOne(context.TODO(), message)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return id.InsertedID.(primitive.ObjectID), nil
}

func (c *ChatRepository) ReadMessage(chatId primitive.ObjectID, senderId uint) (int64, error) {

	filter := bson.M{"chat_id": chatId, "sender_id": senderId, "seen": false}

	update := bson.M{"$set": bson.M{"seen": true}}

	res, err := c.MessageCollection.UpdateMany(context.TODO(), filter, update)

	if err != nil {

		return 0, err
	}

	return res.UpsertedCount, nil

}

func (c *ChatRepository) FetchRecipient(chatId primitive.ObjectID, userId uint) (uint, error) {
	filter := bson.M{"_id": chatId}
	projection := bson.M{"_id": 0, "users": bson.M{"$elemMatch": bson.M{"$ne": userId}}, "created_at": 0}

	chat := domain.Chats{}
	c.ChatCollection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&chat)

	return chat.Users[0], nil

}
