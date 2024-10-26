package http_routes

import "github.com/gorilla/mux"

func SetRoutes(router *mux.Router) {
	SetUsersRoutes(router)
}
