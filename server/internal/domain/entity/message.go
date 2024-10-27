package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ChatID     string             `json:"chatId" bson:"chatId"`
	Chat       Chat               `json:"chat" bson:"-"`
	ReceiverID string             `json:"receiverId" bson:"receiverId"`
	Receiver   User               `json:"receiver" bson:"-"`
	SenderID   string             `json:"senderId" bson:"senderId"`
	Sender     User               `json:"sender" bson:"-"`
	Content    string             `json:"content" bson:"content"`
}
