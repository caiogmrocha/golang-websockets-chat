package infra

import (
	"context"
	"log"
	"os"

	"github.com/caiogmrocha/golang-websockets-chat/server/configs"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoUsersRepository struct {}

func (repo *MongoUsersRepository) GetByEmail(email string) (*entity.User, error) {
  defer func() {
    if r := recover(); r != nil {
      log.Println("Recovered in GetByEmail", r)
    }
  }()

  coll := configs.MongoClient.Database(os.Getenv("MONGO_DB")).Collection("users")

  var user entity.User

  err := coll.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

  if err != nil {
    if err.Error() == "mongo: no documents in result" {
      return nil, nil
    }

    return nil, err
  }

  return &user, nil
}

func (repo *MongoUsersRepository) Create(user *entity.User) error {
  coll := configs.MongoClient.Database(os.Getenv("MONGO_DB")).Collection("users")

  _, err := coll.InsertOne(context.TODO(), user)

  return err
}
