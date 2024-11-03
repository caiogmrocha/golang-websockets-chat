package service

import (
	"errors"
	"time"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUserLastLoginDateService struct {
	UsersRepository repository.UserRepository
}

func (service *UpdateUserLastLoginDateService) Update(userID string) error {
	user, err := service.UsersRepository.GetByID(userID)

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("USER NOT FOUND")
	}

	user.LastLoginDate = primitive.NewDateTimeFromTime(time.Now().Local())

	return service.UsersRepository.Update(user)
}
