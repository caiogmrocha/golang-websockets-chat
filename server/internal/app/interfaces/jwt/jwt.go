package jwt

type JWTProvider interface {
  GenerateToken(payload []byte) (string, error)
  ValidateToken(token string) ([]byte, error)
}
