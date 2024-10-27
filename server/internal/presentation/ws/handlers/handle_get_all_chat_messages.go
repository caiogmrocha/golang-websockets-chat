package ws_handlers

import (
	"encoding/json"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/service"
	infra_repository "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/repository"
	"github.com/olahol/melody"
)

type GetAllChatMessagesHandler struct {
	getAllChatMessagesService *service.GetAllChatMessagesService
}

func (h *GetAllChatMessagesHandler) Handle(s *melody.Session, m *melody.Melody, payload map[string]interface{}) {
	senderID, _ := s.Get("user_id")
	receiverID := payload["receiver_id"].(string)

	messages, err := h.getAllChatMessagesService.Get(senderID.(string), receiverID)

	if err != nil {
		return
	}

	responsePayload := map[string]interface{}{
		"type":     "all_messages",
		"messages": messages,
	}

	marshalledPayload, _ := json.Marshal(responsePayload)

	s.Write(marshalledPayload)
}

func NewGetAllChatMessagesHandler() *GetAllChatMessagesHandler {
	return &GetAllChatMessagesHandler{
		getAllChatMessagesService: &service.GetAllChatMessagesService{
			MessagesRepository: &infra_repository.MongoMessagesRepository{},
		},
	}
}
