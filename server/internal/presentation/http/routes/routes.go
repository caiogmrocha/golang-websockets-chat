package http_routes

import (
	"github.com/gorilla/mux"
	"github.com/olahol/melody"
)

func SetRoutes(router *mux.Router, m *melody.Melody) {
	SetUsersRoutes(router)
	SetWSRoutes(router, m)
}
