package http_routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olahol/melody"
)

func SetWSRoutes(router *mux.Router, m *melody.Melody) {
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})
}
