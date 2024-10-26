package http_controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caiogmrocha/golang-websockets-chat/server/internal/app/service"
	infra_jwt "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/jwt"
	infra_repository "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/repository"
	infra "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/validator"
)

type AuthenticateUserController struct {
	AuthenticateUserService service.AuthenticateUserService
}

func (controller *AuthenticateUserController) Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestBody := make([]byte, r.ContentLength)

	r.Body.Read(requestBody)

	dto := new(service.AuthenticateUserServiceParamsDTO)

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
		case "ERROR_FETCHING_USER":
			{
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "Error fetching user"}`))
			}

		case "USER_NOT_FOUND":
			{
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"error": "User not found"}`))
			}

		case "ERROR_GENERATING_TOKEN":
			{
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "Error generating token"}`))
			}

		default:
			{
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
	mongoUsersRepository := infra_repository.MongoUsersRepository{}
	jwtProvider := infra_jwt.JWTProvider{}

	service := service.AuthenticateUserService{
		UserRepository: &mongoUsersRepository,
		JWTProvider:    &jwtProvider,
	}

	return &AuthenticateUserController{AuthenticateUserService: service}
}
