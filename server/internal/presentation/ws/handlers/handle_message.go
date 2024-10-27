package ws_handlers

import (
	"encoding/json"
	"fmt"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/service"
	infra_repository "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/repository"
	"github.com/olahol/melody"
)

type UserMessageHandler struct {
	registerMessageService *service.RegisterMessageService
}

func (h *UserMessageHandler) HandleMessage(s *melody.Session, m *melody.Melody, payload map[string]interface{}) {
	sessions, err := m.Sessions()

	if err != nil {
		fmt.Println(err)

		return
	}

	for _, session := range sessions {
		if receiverId, exists := session.Get("user_id"); exists && receiverId == payload["receiver_id"] {
			senderID, _ := s.Get("user_id")

			h.registerMessageService.Create(&service.RegisterMessageServiceDTO{
				ReceiverID: payload["receiver_id"].(string),
				SenderID:   senderID.(string),
				Content:    payload["message"].(string),
			})

			responsePayload := map[string]interface{}{
				"type":      "message",
				"sender_id": senderID.(string),
				"message":   payload["message"],
			}

			marshalledPayload, _ := json.Marshal(responsePayload)

			session.Write(marshalledPayload)

			return
		}
	}
}

func NewUserMessageHandler() *UserMessageHandler {
	return &UserMessageHandler{
		registerMessageService: &service.RegisterMessageService{
			MessagesRepository: &infra_repository.MongoMessagesRepository{},
			ChatsRepository:    &infra_repository.MongoChatsRepository{},
		},
	}
}
