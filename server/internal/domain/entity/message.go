package entity

type Message struct {
	ID int
	ChatID int
	Chat Chat
	ReceiverId int
	Receiver User
	SenderId int
	Sender User
	Text string
}