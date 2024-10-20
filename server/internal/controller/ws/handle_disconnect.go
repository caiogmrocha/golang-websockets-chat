package ws

import (
	"encoding/json"
	"log"

	"github.com/olahol/melody"
)

func HandleDisconnect(s *melody.Session, m *melody.Melody) {
	log.Println("Disconnected")

	userId, _ := s.Get("user_id")

	marshalledPayload, _ := json.Marshal(map[string]interface{}{
		"type": "another_user_disconnected",
		"user_id": userId,
	})

	m.BroadcastFilter(marshalledPayload, func (q *melody.Session) bool {
		sId, _ := q.Get("user_id")
		qId, _ := s.Get("user_id")

		return sId != qId
	})
}