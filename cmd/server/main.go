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
	"github.com/chunnior/api-gateway/pkg/infrastructure/logger"
)

func main() {
	logger, err := logger.NewZapLogger()
	if err != nil {
		panic(err)
	}
	// Load environment variables
	err = godotenv.Load()
	if err != nil && os.Getenv("ENV") == "local" {
		logger.Error("Error loading .env file", "error", err)
	}

	// Create a new Fiber instance
	app := fiber.New()

	// Setup middleware
	middleware.SetupMiddleware(app)

	// pemFileName := "public_key.pem"

	// // Leer la clave pública desde el archivo PEM
	// publicKey, err := utils.ReadPublicKeyFromFile(pemFileName)
	// if err != nil {
	// 	fmt.Println("Error al leer la clave pública desde el archivo PEM:", err)
	// 	return
	// }

	//
	/*
		producer, err := kafka.NewProducer(os.Getenv("KAFKA_HOST"))
		if err != nil {
			log.Fatalf("Error al crear el productor: %v", err)
		}
		defer producer.Close()
	*/

	// Inicializa el servicios
	userService := infrastructure.NewHTTPUserService(os.Getenv("USER_SERVICE_URL"), &http.Client{}, os.Getenv("API_KEY"))
	authService := domain.NewAuthServiceImpl(os.Getenv("JWT_SECRET_KEY"))

	// Initialize handlers
	healthHandler := handler.NewHealthHandler()
	userHandler := handler.NewUserHandler(userService, authService)

	// authService := domain.NewAuthServiceImpl(os.Getenv("JWT_SECRET_KEY"))
	jwtMiddleware := middleware.NewJWTMiddleware(authService, logger)

	// Setup routes
	router.SetupRoutes(app, jwtMiddleware, userHandler, healthHandler)

	// Start the server
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
