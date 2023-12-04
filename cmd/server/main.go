package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/chunnior/api-gateway/internal/app/handler"
	"github.com/chunnior/api-gateway/internal/app/middleware"
	"github.com/chunnior/api-gateway/internal/app/router"
	"github.com/chunnior/api-gateway/internal/domain"

	// "github.com/chunnior/api-gateway/internal/repository"
	// "github.com/chunnior/api-gateway/internal/usecase"
	"github.com/chunnior/api-gateway/pkg/infrastructure"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new Fiber instance
	app := fiber.New()

	// Setup middleware
	middleware.SetupMiddleware(app)

	// Inicializa el servicios
	userService := infrastructure.NewHTTPUserService(os.Getenv("USER_SERVICE_URL"), &http.Client{})
	authService := domain.NewAuthServiceImpl(os.Getenv("JWT_SECRET_KEY"))

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService, authService)

	// authService := domain.NewAuthServiceImpl(os.Getenv("JWT_SECRET_KEY"))
	jwtMiddleware := middleware.NewJWTMiddleware(authService)

	// Setup routes
	router.SetupRoutes(app, jwtMiddleware, userHandler)

	// Start the server
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
