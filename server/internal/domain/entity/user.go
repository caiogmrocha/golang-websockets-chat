package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name"`
	Email         string             `json:"email" bson:"email"`
	Password      string             `json:"password" bson:"password"`
	LastLoginDate primitive.DateTime `json:"last_login_date" bson:"last_login_date,omitempty"`
}
