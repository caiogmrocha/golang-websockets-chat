package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	services "github.com/caiogmrocha/golang-websockets-chat/server/internal/app/service"
	jwt_impl "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/jwt"
	repositories_impl "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/repositories/infra"
	infra "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/validator"
)

type AuthenticateUserController struct {
  AuthenticateUserService services.AuthenticateUserService
}

func (controller *AuthenticateUserController) Authenticate(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  requestBody := make([]byte, r.ContentLength)

  r.Body.Read(requestBody)

  dto := new(services.AuthenticateUserServiceParamsDTO)

  err := json.Unmarshal(requestBody, dto)

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(`{"error": "Invalid JSON"}`))

    return
  }

  customValidatorErr := infra.CustomValidator(dto)

  if customValidatorErr != nil {
    marshalledErr, _ := json.Marshal(customValidatorErr)

    w.WriteHeader(http.StatusBadRequest)
    w.Write(marshalledErr)

    return
  }

  token, err := controller.AuthenticateUserService.Authenticate(dto)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)

    switch err.Error() {
      case "ERROR_FETCHING_USER": {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"error": "Error fetching user"}`))
      }

      case "USER_NOT_FOUND": {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"error": "User not found"}`))
      }

      case "ERROR_GENERATING_TOKEN": {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"error": "Error generating token"}`))
      }

      default: {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"error": "Internal Server Error"}`))
      }
    }

    return
  }

  if token == "" {
    w.WriteHeader(http.StatusUnauthorized)
    w.Write([]byte(`{"error": "Unauthorized"}`))

    return
  }

  w.WriteHeader(http.StatusOK)
  w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
}

func NewAuthenticateUserController() *AuthenticateUserController {
  mongoUsersRepository := repositories_impl.MongoUsersRepository{}
  jwtProvider := jwt_impl.JWTProvider{}

  service := services.AuthenticateUserService{
    UserRepository: &mongoUsersRepository,
    JWTProvider: &jwtProvider,
  }

  return &AuthenticateUserController{AuthenticateUserService: service}
}
