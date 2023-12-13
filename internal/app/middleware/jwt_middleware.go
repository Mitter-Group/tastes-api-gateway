package middleware

import (
	"github.com/chunnior/api-gateway/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type JWTMiddleware struct {
	authService domain.AuthService
}

func NewJWTMiddleware(authService domain.AuthService) *JWTMiddleware {
	return &JWTMiddleware{
		authService: authService,
	}
}

func (m *JWTMiddleware) ValidateJWT(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Authorization token missing",
		})
	}

	user, err := m.authService.ValidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	c.Locals("user", user)

	return c.Next()
}
