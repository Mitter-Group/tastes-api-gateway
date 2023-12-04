package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// SetupMiddleware sets up the middleware for the application.
func SetupMiddleware(app *fiber.App) {
	// Add recovery middleware to recover from panics
	app.Use(recover.New())

	// Add logger middleware to log HTTP requests
	app.Use(logger.New())
}