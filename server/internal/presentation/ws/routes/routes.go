package ws_routes

import (
	"encoding/json"

	ws_handlers "github.com/caiogmrocha/golang-websockets-chat/server/internal/presentation/ws/handlers"
	"github.com/olahol/melody"
)

func SetWSHandlers(m *melody.Melody) {
	connectionHandler := ws_handlers.NewConnectHandler()
	m.HandleConnect(func(s *melody.Session) { connectionHandler.HandleConnect(s, m) })

	m.HandleDisconnect(func(s *melody.Session) { ws_handlers.HandleDisconnect(s, m) })

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		var payload map[string]interface{}

		json.Unmarshal(msg, &payload)

		switch payload["type"] {
		case "message":
			UserMessageHandler := ws_handlers.NewUserMessageHandler()

			UserMessageHandler.HandleMessage(s, m, payload)
		case "connected_users":
			GetUsersHandler := ws_handlers.NewGetUsersHandler()

			GetUsersHandler.HandleGetUsers(s, m, payload)
		case "all_messages":
			GetAllChatMessagesHandler := ws_handlers.NewGetAllChatMessagesHandler()

			GetAllChatMessagesHandler.Handle(s, m, payload)
		}
	})
}
