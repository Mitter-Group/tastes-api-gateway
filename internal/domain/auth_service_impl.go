package domain

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthServiceImpl struct {
	secretKey string
}

type Claims struct {
	UserPayload
	jwt.StandardClaims
}

func NewAuthServiceImpl(secretKey string) *AuthServiceImpl {
	return &AuthServiceImpl{
		secretKey: secretKey,
	}
}

func (s *AuthServiceImpl) ValidateToken(tokenString string) (userID string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Asegúrate de que el algoritmo de token es lo que esperas
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Retorna tu clave secreta
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Puedes acceder a los claims del token aquí, por ejemplo:
		userID, ok := claims["sub"].(string)
		if !ok {
			return "", fmt.Errorf("error getting userID from token")
		}

		return userID, nil
	}

	return "", fmt.Errorf("invalid token")
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
