package ws_handlers

import (
	"encoding/json"
	"fmt"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/service"
	infra_repository "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/repository"
	"github.com/olahol/melody"
)

type GetUsersHandler struct {
	GetUsersByIdService *service.GetUsersByIdService
}

func (h *GetUsersHandler) HandleGetUsers(s *melody.Session, m *melody.Melody, payload map[string]interface{}) {
	sessions, err := m.Sessions()

	if err != nil {
		fmt.Println(err)

		return
	}

	usersIds := []string{}

	for _, session := range sessions {
		if userId, exists := session.Get("user_id"); exists {
			usersIds = append(usersIds, userId.(string))
		}
	}

	users, error := h.GetUsersByIdService.Get(usersIds)

	if error != nil {
		fmt.Println(error)

		return
	}

	responsePayload := map[string]interface{}{
		"type":  "connected_users",
		"users": users,
	}

	marshalledPayload, _ := json.Marshal(responsePayload)

	s.Write(marshalledPayload)
}

func NewGetUsersHandler() *GetUsersHandler {
	usersRepository := infra_repository.MongoUsersRepository{}
	getUsersByIdService := service.GetUsersByIdService{
		UsersRepository: &usersRepository,
	}

	return &GetUsersHandler{
		GetUsersByIdService: &getUsersByIdService,
	}
}
