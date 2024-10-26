package main

import (
	"context"
	"log"
	"net/http"

	"github.com/caiogmrocha/golang-websockets-chat/server/configs"
	http_routes "github.com/caiogmrocha/golang-websockets-chat/server/internal/presentation/http/routes"
	ws_routes "github.com/caiogmrocha/golang-websockets-chat/server/internal/presentation/ws/routes"
	"github.com/gorilla/mux"
	"github.com/olahol/melody"
)

func main() {
	defer configs.MongoClient.Disconnect(context.Background())

	router := mux.NewRouter()

	m := melody.New()

	defer m.Close()

	http_routes.SetRoutes(router, m)
	ws_routes.SetWSHandlers(m)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
