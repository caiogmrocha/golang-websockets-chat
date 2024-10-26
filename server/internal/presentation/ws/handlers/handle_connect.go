package ws_handlers

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/olahol/melody"
)

func HandleConnect(s *melody.Session, m *melody.Melody) {
	log.Println(s.Request.Header.Get("User-ID"))
	log.Println("Connected")

	userId := uuid.New().String()

	s.Set("user_id", userId)

	responsePayload := map[string]interface{}{
		"type":    "user_id",
		"user_id": userId,
	}

	marshalledPayload, _ := json.Marshal(responsePayload)

	s.Write(marshalledPayload)

	marshalledPayload, _ = json.Marshal(map[string]interface{}{
		"type":    "another_user_connected",
		"user_id": userId,
	})

	m.BroadcastFilter(marshalledPayload, func(q *melody.Session) bool {
		sId, _ := q.Get("user_id")
		qId, _ := s.Get("user_id")

		return sId != qId
	})
}
