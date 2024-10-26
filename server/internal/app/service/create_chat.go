package service

import "github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repository"

type CreateChatServiceDTO struct {
	UsersIDs []int `json:"users_id"`
}

type CreateChatService struct {
	ChatRepository repository.ChatRepository
}
