package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/olahol/melody"
)

func main() {
	m := melody.New()

	http.HandleFunc("/ws", func (w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	// users := map[string]string{}
	chats := map[string][]string{}

	m.HandleMessage(func (s *melody.Session, msg []byte) {
		userId := s.Request.Header.Get("User-ID")
		
		var clientPayload map[string]interface{}

		json.Unmarshal(msg, &clientPayload)
		
		chatKey := fmt.Sprintf("%s-%s", userId, clientPayload["receiver_id"])

		s.Set(chatKey, true)

		switch clientPayload["type"] {
			case "sendMessage": {
				sessions, err := m.Sessions()

				if err != nil {
					fmt.Println(err)

					return
				}

				chats[chatKey] = append(chats[chatKey], clientPayload["message"].(string))

				for _, session := range sessions {
					if session.Request.Header.Get("User-ID") == clientPayload["receiver_id"] {
						responsePayload := map[string]interface{}{
							"type": "message",
							"sender_id": userId,
							"message": clientPayload["message"],
						}

						marshalledPayload, _ := json.Marshal(responsePayload)

						session.Write(marshalledPayload)
					}
				}
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}