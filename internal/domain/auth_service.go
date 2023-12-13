package domain

type AuthService interface {
	ValidateToken(token string) (*UserPayload, error)
	GenerateTokens(userPayload UserPayload) (string, string, error)
}
