package infra_repository

import (
	"context"
	"os"

	"github.com/caiogmrocha/golang-websockets-chat/server/configs"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"
)

type MongoMessagesRepository struct{}

func (repo *MongoMessagesRepository) Create(message *entity.Message) error {
	coll := configs.MongoClient.Database(os.Getenv("MONGO_DB")).Collection("messages")

	_, err := coll.InsertOne(context.TODO(), message)

	return err
}
