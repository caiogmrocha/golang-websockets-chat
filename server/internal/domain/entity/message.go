package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ChatID     string             `json:"chat_id" bson:"chatId"`
	Chat       Chat               `json:"chat" bson:"-"`
	ReceiverId string             `json:"receiver_id" bson:"receiverId"`
	Receiver   User               `json:"receiver" bson:"-"`
	SenderId   string             `json:"sender_id" bson:"senderId"`
	Sender     User               `json:"sender" bson:"-"`
	Content    string             `json:"content" bson:"content"`
}
