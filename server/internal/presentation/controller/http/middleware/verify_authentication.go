package http_middleware

import (
	"context"
	"encoding/json"
	"net/http"

	infra_jwt "github.com/caiogmrocha/golang-websockets-chat/server/internal/infra/jwt"
)

type VerifyAuthenticationHTTPMiddleware struct {
  JWTProvider infra_jwt.JWTProvider
}

type key string

func (m *VerifyAuthenticationHTTPMiddleware) Handle(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    token := r.Header.Get("Authorization")

    if token == "" {
      http.Error(w, "Token not provided", http.StatusUnauthorized)
      return
    }

    payload, err := m.JWTProvider.ValidateToken(token)

    if err != nil {
      http.Error(w, "Invalid token", http.StatusUnauthorized)
      return
    }

    var unmarsheledPayload map[string]interface{}

    if err := json.Unmarshal(payload, &unmarsheledPayload); err != nil {
      http.Error(w, "Error unmarshalling token payload", http.StatusInternalServerError)
      return
    }

    ctx := context.WithValue(r.Context(), key("userID"), unmarsheledPayload["sub"])

    next.ServeHTTP(w, r.WithContext(ctx))
  })
}
