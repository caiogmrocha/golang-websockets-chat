package http_middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	infra_jwt "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/jwt"
)

type VerifyAuthenticationHTTPMiddleware struct {
	JWTProvider infra_jwt.JWTProvider
}

type Key string

func (m *VerifyAuthenticationHTTPMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware")

		tokenCookie, err := r.Cookie("token")

		if err != nil {
			log.Println("Middleware err", err)

			http.Error(w, "Token not provided", http.StatusUnauthorized)
			return
		}

		log.Println("Middleware tokenCookie.Value", tokenCookie.Value)

		if tokenCookie.Value == "" {
			http.Error(w, "Token not provided", http.StatusUnauthorized)
			return
		}

		payload, err := m.JWTProvider.ValidateToken(tokenCookie.Value)

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		var unmarsheledPayload map[string]interface{}

		if err := json.Unmarshal(payload, &unmarsheledPayload); err != nil {
			http.Error(w, "Error unmarshalling token payload", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), Key("userID"), unmarsheledPayload["id"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewVerifyAuthenticationHTTPMiddleware() *VerifyAuthenticationHTTPMiddleware {
	return &VerifyAuthenticationHTTPMiddleware{
		JWTProvider: infra_jwt.JWTProvider{},
	}
}
