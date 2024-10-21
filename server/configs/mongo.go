package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func init() {
  uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
    os.Getenv("MONGO_USER"),
    os.Getenv("MONGO_PASS"),
    os.Getenv("MONGO_HOST"),
    os.Getenv("MONGO_PORT"),
  )

  MongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  err = MongoClient.Ping(context.Background(), nil)

  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  log.Println("Connected to MongoDB")
}
