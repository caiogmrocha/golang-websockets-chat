package infra_repository

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/caiogmrocha/golang-websockets-chat/server/configs"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"
)

type MongoMessagesRepository struct{}

func (repo *MongoMessagesRepository) GetBySenderIdAndReceiverId(senderID, receiverID string) ([]entity.Message, error) {
	coll := configs.MongoClient.Database(os.Getenv("MONGO_DB")).Collection("messages")

	cursor, err := coll.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"senderId": senderID, "receiverId": receiverID},
			{"senderId": receiverID, "receiverId": senderID},
		},
	})

	if err != nil {
		return nil, err
	}

	messages := []entity.Message{}

	err = cursor.All(context.TODO(), &messages)

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (repo *MongoMessagesRepository) GetByChatID(chatID string) ([]entity.Message, error) {
	coll := configs.MongoClient.Database(os.Getenv("MONGO_DB")).Collection("messages")

	cursor, err := coll.Find(context.TODO(), bson.M{"chatId": chatID})

	if err != nil {
		return nil, err
	}

	var messages []entity.Message

	err = cursor.All(context.TODO(), &messages)

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (repo *MongoMessagesRepository) Create(message *entity.Message) error {
	coll := configs.MongoClient.Database(os.Getenv("MONGO_DB")).Collection("messages")

	_, err := coll.InsertOne(context.TODO(), message)

	return err
}
