package http_routes

import (
	"net/http"

	http_middleware "github.com/caiogmrocha/golang-websockets-chat/server/internal/presentation/http/middleware"
	"github.com/gorilla/mux"
	"github.com/olahol/melody"
)

func SetWSRoutes(router *mux.Router, m *melody.Melody) {
	subrouter := router.NewRoute().Subrouter()

	verifyAuthenticationHTTPMiddleware := http_middleware.NewVerifyAuthenticationHTTPMiddleware()

	subrouter.Use(verifyAuthenticationHTTPMiddleware.Handle)

	subrouter.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(http_middleware.Key("userID"))

		m.HandleRequestWithKeys(w, r, map[string]interface{}{
			"user_id": userID,
		})
	})
}
