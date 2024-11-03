package infra_repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/caiogmrocha/golang-websockets-chat/server/configs"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoUsersRepository struct{}

func (repo *MongoUsersRepository) GetByEmail(email string) (*entity.User, error) {
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

func (repo *MongoUsersRepository) GetByID(id string) (*entity.User, error) {
	coll := configs.MongoClient.Database(os.Getenv("MONGO_DB")).Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var user entity.User

	fmt.Print("Aquele ID lascado: ")
	fmt.Println(id)

	err = coll.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)

	fmt.Print("Aquele erro lascado: ")
	fmt.Println(err)
	fmt.Print("Aquele user lascado: ")
	fmt.Println(user)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (repo *MongoUsersRepository) GetManyById(ids []string) ([]*entity.User, error) {
	coll := configs.MongoClient.Database(os.Getenv("MONGO_DB")).Collection("users")

	users := []*entity.User{}

	objectsIds := []primitive.ObjectID{}

	for _, id := range ids {
		objectID, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return nil, err
		}

		objectsIds = append(objectsIds, objectID)
	}

	cursor, err := coll.Find(context.TODO(), bson.M{"_id": bson.M{"$in": objectsIds}})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user entity.User

		err := cursor.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (repo *MongoUsersRepository) Create(user *entity.User) error {
	coll := configs.MongoClient.Database(os.Getenv("MONGO_DB")).Collection("users")

	_, err := coll.InsertOne(context.TODO(), user)

	return err
}

func (repo *MongoUsersRepository) Update(user *entity.User) error {
	coll := configs.MongoClient.Database(os.Getenv("MONGO_DB")).Collection("users")

	_, err := coll.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.M{"$set": user})

	return err
}

func (repo *MongoUsersRepository) DeleteInactiveUsers(bound time.Time) error {
	coll := configs.MongoClient.Database(os.Getenv("MONGO_DB")).Collection("users")

	_, err := coll.DeleteMany(context.TODO(), bson.M{"last_login_date": bson.M{"$lt": bound}})

	return err
}
