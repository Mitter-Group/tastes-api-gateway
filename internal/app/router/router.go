package router

import (
	"github.com/chunnior/api-gateway/internal/app/handler"
	"github.com/chunnior/api-gateway/internal/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, jwtMiddleware *middleware.JWTMiddleware, userHandler *handler.UserHandler) {
	api := app.Group("/api")

	usersRoutes := api.Group("/users")
	usersRoutes.Post("/login", userHandler.Login)
	usersRoutes.Get("/callback", userHandler.Callback)
	usersRoutes.Post("/refresh-token", userHandler.RefreshToken)
	usersRoutes.Get("/:provider/:dataType", jwtMiddleware.ValidateJWT, userHandler.HandleDataInfo)
}
