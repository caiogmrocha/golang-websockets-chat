package repository

import (
	"time"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"
)

type UserRepository interface {
	GetByID(id string) (*entity.User, error)
	GetManyById(ids []string) ([]*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	DeleteInactiveUsers(bound time.Time) error
}
