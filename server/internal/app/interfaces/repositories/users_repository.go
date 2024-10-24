package repositories

import "github.com/caiogmrocha/golang-websockets-chat/server/internal/domain/entity"

type UserRepository interface {
  GetByEmail(email string) (*entity.User, error)
  Create(user *entity.User) error
}
