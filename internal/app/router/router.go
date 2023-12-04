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

	// usersRoutes.Get("/", userHandler.GetUsers)
	/*
		usersRoutes.Get("/:id", jwtMiddleware.ValidateJWT, userHandler.GetUser)
		usersRoutes.Post("/", userHandler.CreateUser)
	*/
	// usersRoutes.Put("/:id", userHandler.UpdateUser)
	// usersRoutes.Delete("/:id", userHandler.DeleteUser)

	/*
		geoRoutes := api.Group("/geo")
		// http.Handle("/products", ValidateJWT(http.HandlerFunc(YourHandlerFunc)))
		api.Get("/products", productHandler.GetProducts)
		api.Get("/products/:id", productHandler.GetProduct)
		api.Post("/products", productHandler.CreateProduct)
		api.Put("/products/:id", productHandler.UpdateProduct)
		api.Delete("/products/:id", productHandler.DeleteProduct)
	*/
}
