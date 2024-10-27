package repository

import "github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"

type ChatsRepository interface {
	GetByUsersIDs(usersIDs [2]string) (*entity.Chat, error)
	Create(chat *entity.Chat) error
}
