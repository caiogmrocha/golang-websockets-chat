package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/olahol/melody"

	"github.com/caiogmrocha/golang-websockets-chat/server/configs"
	http_routes "github.com/caiogmrocha/golang-websockets-chat/server/internal/presentation/http/routes"
	ws_routes "github.com/caiogmrocha/golang-websockets-chat/server/internal/presentation/ws/routes"
)

func main() {
	defer configs.MongoClient.Disconnect(context.Background())

	router := mux.NewRouter()

	m := melody.New()

	defer m.Close()

	http_routes.SetRoutes(router, m)
	ws_routes.SetWSHandlers(m)

	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	cors := handlers.CORS(
		handlers.AllowedOrigins(allowedOrigins),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", cors(router)))
}
