package http_routes

import (
	http_controller "github.com/caiogmrocha/golang-websockets-chat/server/internal/presentation/http/controller"
	"github.com/gorilla/mux"
)

func SetUsersRoutes(router *mux.Router) {
	registerUserController := http_controller.NewRegisterUserController()
	authenticateUserController := http_controller.NewAuthenticateUserController()

	router.HandleFunc("/users", registerUserController.Create).Methods("POST").Headers("Content-Type", "application/json")
	router.HandleFunc("/users/authenticate", authenticateUserController.Authenticate).Methods("POST").Headers("Content-Type", "application/json")
}
