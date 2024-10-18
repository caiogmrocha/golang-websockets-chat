package entity

type Chat struct {
	ID int
	Name string
	Users []User
	Messages []Message
}