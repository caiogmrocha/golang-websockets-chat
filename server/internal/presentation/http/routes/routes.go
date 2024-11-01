package http_routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olahol/melody"
)

func SetRoutes(router *mux.Router, m *melody.Melody) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"status\": \"ok\"}"))
	}).Methods("GET")

	SetUsersRoutes(router)
	SetWSRoutes(router, m)
}
