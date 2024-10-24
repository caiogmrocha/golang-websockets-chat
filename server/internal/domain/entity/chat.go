package entity

type Chat struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Users [2]User `json:"users"`
	Messages []Message `json:"messages"`
}
