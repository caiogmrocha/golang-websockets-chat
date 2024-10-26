package entity

type Chat struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	UsersIDs    [2]string `json:"users_ids" bson:"usersIds"`
	Users       [2]User   `json:"users" bson:"-"`
	MessagesIDs []string  `json:"messages_ids" bson:"messagesIds"`
	Messages    []Message `json:"messages" bson:"-"`
}
