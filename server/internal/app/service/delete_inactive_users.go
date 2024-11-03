package service

import (
	"time"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/interfaces/repository"
	infra_repository "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/repository"
)

type DeleteInactiveUsersService struct {
	UsersRepository repository.UserRepository
}

func (service *DeleteInactiveUsersService) DeleteInactiveUsers() error {
	bound := time.Now().AddDate(0, 0, -7)

	return service.UsersRepository.DeleteInactiveUsers(bound)
}

func NewDeleteInactiveUsersService() *DeleteInactiveUsersService {
	usersRepository := infra_repository.MongoUsersRepository{}
	return &DeleteInactiveUsersService{
		UsersRepository: &usersRepository,
	}
}
