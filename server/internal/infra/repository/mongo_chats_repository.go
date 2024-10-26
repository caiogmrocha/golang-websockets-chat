package infra_repository

import (
	"context"
	"os"

	"github.com/caiogmrocha/golang-websockets-chat/server/configs"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoChatsRepository struct{}

func (repo *MongoChatsRepository) GetByUsersIDs(usersIDs [2]string) (*entity.Chat, error) {
	coll := configs.MongoClient.Database(os.Getenv("MONGO_DB")).Collection("chats")

	var chat entity.Chat

	err := coll.FindOne(context.TODO(), bson.M{"users_ids": usersIDs}).Decode(&chat)

	if err == nil {
		return &chat, nil
	}

	if err.Error() != "mongo: no documents in result" {
		return nil, err
	}

	usersIDs = [2]string{usersIDs[1], usersIDs[0]}

	err = coll.FindOne(context.TODO(), bson.M{"users_ids": usersIDs}).Decode(&chat)

	if err == nil {
		return &chat, nil
	}

	if err.Error() == "mongo: no documents in result" {
		return nil, nil
	}

	return nil, err
}
