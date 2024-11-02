package http_routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olahol/melody"
)

func SetRoutes(router *mux.Router, m *melody.Melody) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		responseBody, _ := json.Marshal(map[string]string{"status": "ok"})

		w.Write(responseBody)
	}).Methods("GET")

	SetUsersRoutes(router)
	SetWSRoutes(router, m)
}
