package ws_routes

import (
	"encoding/json"

	ws_handlers "github.com/caiogmrocha/golang-websockets-chat/server/internal/presentation/ws/handlers"
	"github.com/olahol/melody"
)

func SetWSHandlers(m *melody.Melody) {
	m.HandleConnect(func(s *melody.Session) { ws_handlers.HandleConnect(s, m) })
	m.HandleDisconnect(func(s *melody.Session) { ws_handlers.HandleDisconnect(s, m) })

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		var payload map[string]interface{}

		json.Unmarshal(msg, &payload)

		switch payload["type"] {
		case "message":
			UserMessageHandler := ws_handlers.NewUserMessageHandler()

			UserMessageHandler.HandleMessage(s, m, payload)
		case "users_ids":
			ws_handlers.HandleGetUsersIds(s, m)
		case "all_messages":
			GetAllChatMessagesHandler := ws_handlers.NewGetAllChatMessagesHandler()

			GetAllChatMessagesHandler.Handle(s, m, payload)
		}
	})
}
