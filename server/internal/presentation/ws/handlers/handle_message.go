package ws_handlers

import (
	"encoding/json"
	"fmt"

	"github.com/olahol/melody"
)

func HandleMessage(payload map[string]interface{}, m *melody.Melody) {
	sessions, err := m.Sessions()

	if err != nil {
		fmt.Println(err)

		return
	}

	for _, session := range sessions {
		if receiverId, exists := session.Get("user_id"); exists && receiverId == payload["receiver_id"] {
			responsePayload := map[string]interface{}{
				"type":      "message",
				"sender_id": payload["sender_id"],
				"message":   payload["message"],
			}

			marshalledPayload, _ := json.Marshal(responsePayload)

			session.Write(marshalledPayload)
		}
	}
}
