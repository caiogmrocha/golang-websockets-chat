package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/olahol/melody"
)

func main() {
	m := melody.New()

	http.HandleFunc("/ws", func (w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	m.HandleConnect(func (s *melody.Session) {
		log.Println("Connected")
		
		userId := uuid.New().String()

		s.Set("user_id", userId)

		responsePayload := map[string]interface{}{
			"type": "user_id",
			"user_id": userId,
		}

		marshalledPayload, _ := json.Marshal(responsePayload)

		s.Write(marshalledPayload)

		marshalledPayload, _ = json.Marshal(map[string]interface{}{
			"type": "another_user_connected",
			"user_id": userId,
		})

		m.BroadcastFilter(marshalledPayload, func (q *melody.Session) bool {
			sId, _ := q.Get("user_id")
			qId, _ := s.Get("user_id")

			return sId != qId
		})
	})

	m.HandleMessage(func (s *melody.Session, msg []byte) {
		var clientPayload map[string]interface{}
		
		json.Unmarshal(msg, &clientPayload)

		switch clientPayload["type"] {
			case "message": {
				sessions, err := m.Sessions()

				if err != nil {
					fmt.Println(err)

					return
				}

				for _, session := range sessions {
					if receiverId, exists := session.Get("user_id"); exists && receiverId == clientPayload["receiver_id"] {
						responsePayload := map[string]interface{}{
							"type": "message",
							"sender_id": clientPayload["sender_id"],
							"message": clientPayload["message"],
						}

						marshalledPayload, _ := json.Marshal(responsePayload)

						session.Write(marshalledPayload)
					}
				}
			}

			case "users_ids": {
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
					"type": "users_ids",
					"users_ids": usersIds,
				}

				marshalledPayload, _ := json.Marshal(responsePayload)

				s.Write(marshalledPayload)
			}
		}


	})

	log.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}