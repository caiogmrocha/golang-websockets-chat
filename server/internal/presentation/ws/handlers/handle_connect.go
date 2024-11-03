package ws_handlers

import (
	"encoding/json"
	"log"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/service"
	infra_repository "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/repository"
	"github.com/olahol/melody"
)

type ConnectHandler struct {
	GetUserByIdService *service.GetUserByIdService
}

func (h *ConnectHandler) HandleConnect(s *melody.Session, m *melody.Melody) {
	userID, _ := s.Get("user_id")

	log.Println("Connected")

	responsePayload := map[string]interface{}{
		"type":    "user_id",
		"user_id": userID,
	}

	marshalledPayload, _ := json.Marshal(responsePayload)

	s.Write(marshalledPayload)

	user, err := h.GetUserByIdService.Get(userID.(string))

	if err != nil {
		log.Println(err)

		return
	}

	marshalledPayload, _ = json.Marshal(map[string]interface{}{
		"type": "another_user_connected",
		"user": user,
	})

	m.BroadcastFilter(marshalledPayload, func(q *melody.Session) bool {
		sId, _ := q.Get("user_id")
		qId, _ := s.Get("user_id")

		return sId != qId
	})
}

func NewConnectHandler() *ConnectHandler {
	usersRepository := infra_repository.MongoUsersRepository{}
	getUserByIdService := service.GetUserByIdService{
		UsersRepository: &usersRepository,
	}

	return &ConnectHandler{
		GetUserByIdService: &getUserByIdService,
	}
}
