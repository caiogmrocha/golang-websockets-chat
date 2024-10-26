package entity

type Message struct {
	ID         int    `json:"id"`
	ChatID     int    `json:"chat_id"`
	Chat       Chat   `json:"chat"`
	ReceiverId int    `json:"receiver_id"`
	Receiver   User   `json:"receiver"`
	SenderId   int    `json:"sender_id"`
	Sender     User   `json:"sender"`
	Content    string `json:"content"`
}
