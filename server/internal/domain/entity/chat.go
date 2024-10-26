package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UsersIDs    [2]string          `json:"users_ids" bson:"usersIds"`
	Users       [2]User            `json:"users" bson:"-"`
	MessagesIDs []string           `json:"messages_ids" bson:"messagesIds"`
	Messages    []Message          `json:"messages" bson:"-"`
}
