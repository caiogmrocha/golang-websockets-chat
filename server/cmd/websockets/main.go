package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/caiogmrocha/golang-websockets-chat/server/configs"
	controllers "github.com/caiogmrocha/golang-websockets-chat/server/internal/controller/http"
	"github.com/caiogmrocha/golang-websockets-chat/server/internal/controller/ws"
	"github.com/gorilla/mux"
	"github.com/olahol/melody"
)

func main() {
  defer configs.MongoClient.Disconnect(context.Background())

  // HTTP Routes
  router := mux.NewRouter()

  registerUserController := controllers.NewRegisterUserController()

  router.HandleFunc("/users", registerUserController.Create).Methods("POST").Headers("Content-Type", "application/json")

  // Websockets Handlers
  m := melody.New()

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

  router.HandleFunc("/ws", func (w http.ResponseWriter, r *http.Request) {
    m.HandleRequest(w, r)
  })

  defer m.Close()

  log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
