package http_controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	services "github.com/caiogmrocha/golang-websockets-chat/server/internal/app/service"
	infra_repository "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/repository"
	infra_validator "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/validator"
)

type RegisterUserController struct {
	RegisterUserService services.RegisterUserService
}

func (controller *RegisterUserController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestBody := make([]byte, r.ContentLength)

	r.Body.Read(requestBody)

	dto := new(services.RegisterUserServiceDTO)

	err := json.Unmarshal(requestBody, dto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid JSON"}`))

		return
	}

	customValidatorErr := infra_validator.CustomValidator(dto)

	if customValidatorErr != nil {
		marshalledErr, _ := json.Marshal(customValidatorErr)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(marshalledErr)

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))

		return
	}

	dto.Password = string(hash)

	err = controller.RegisterUserService.Create(dto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User created successfully"}`))
}

func NewRegisterUserController() *RegisterUserController {
	mongoUsersRepository := infra_repository.MongoUsersRepository{}
	service := services.RegisterUserService{UserRepository: &mongoUsersRepository}

	return &RegisterUserController{RegisterUserService: service}
}
