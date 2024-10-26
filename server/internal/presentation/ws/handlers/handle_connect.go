package ws_handlers

import (
	"encoding/json"
	"log"

	"github.com/olahol/melody"
)

func HandleConnect(s *melody.Session, m *melody.Melody) {
	userID := s.Request.Context().Value("userID")

	log.Println(userID)
	log.Println("Connected")

	s.Set("user_id", userID)

	responsePayload := map[string]interface{}{
		"type":    "user_id",
		"user_id": userID,
	}

	marshalledPayload, _ := json.Marshal(responsePayload)

	s.Write(marshalledPayload)

	marshalledPayload, _ = json.Marshal(map[string]interface{}{
		"type":    "another_user_connected",
		"user_id": userID,
	})

	m.BroadcastFilter(marshalledPayload, func(q *melody.Session) bool {
		sId, _ := q.Get("user_id")
		qId, _ := s.Get("user_id")

		return sId != qId
	})
}
