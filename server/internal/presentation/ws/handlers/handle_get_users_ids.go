package ws_handlers

import (
	"encoding/json"
	"fmt"

	"github.com/olahol/melody"
)

func HandleGetUsersIds(s *melody.Session, m *melody.Melody) {
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

	responsePayload := map[string]interface{}{
		"type":      "users_ids",
		"users_ids": usersIds,
	}

	marshalledPayload, _ := json.Marshal(responsePayload)

	s.Write(marshalledPayload)
}
