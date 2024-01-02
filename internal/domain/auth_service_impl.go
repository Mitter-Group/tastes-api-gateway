package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthServiceImpl struct {
	secretKey []byte
}

type Claims struct {
	UserPayload
	jwt.StandardClaims
}

func NewAuthServiceImpl(secretKey string) *AuthServiceImpl {
	return &AuthServiceImpl{
		secretKey: []byte(secretKey),
	}
}

func (s *AuthServiceImpl) ValidateToken(tokenString string) (*UserPayload, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Asegúrate de que el algoritmo de token es lo que esperas
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Retorna tu clave secreta
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userPayload := &UserPayload{
			ID:             claims["ID"].(string),
			Provider:       claims["provider"].(string),
			ProviderUserID: claims["provider_user_id"].(string),
			UserFullname:   claims["user_fullname"].(string),
			Email:          claims["email"].(string),
		}

		return userPayload, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func (s *AuthServiceImpl) GenerateTokens(userPayload UserPayload) (string, string, error) {
	// Generar Access Token
	accessTokenClaims := &Claims{
		UserPayload: userPayload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(), // Expire at 15 minutes
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(s.secretKey)
	if err != nil {
		return "", "", err
	}

	// Generar Refresh Token
	refreshTokenClaims := &Claims{
		UserPayload: userPayload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * 7 * time.Hour).Unix(), // Expire at 7 days
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(s.secretKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
