package domain

type AuthService interface {
	ValidateToken(token string) (userID string, err error)
	GenerateTokens(userPayload UserPayload) (string, string, error)
}
