package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/caiogmrocha/golang-websockets-chat/internal/controller/ws"
	"github.com/olahol/melody"
)

func main() {
	m := melody.New()

	http.HandleFunc("/ws", func (w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

  defer m.Close()

	m.HandleConnect(func (s *melody.Session) { ws.HandleConnect(s, m) })
	m.HandleDisconnect(func (s *melody.Session) { ws.HandleDisconnect(s, m) })

	m.HandleMessage(func (s *melody.Session, msg []byte) {
		var payload map[string]interface{}

		json.Unmarshal(msg, &payload)

		switch payload["type"] {
			case "message": ws.HandleMessage(payload, m)
			case "users_ids": ws.HandleGetUsersIds(s, m)
		}
	})

	log.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
