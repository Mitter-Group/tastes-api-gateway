package handler

import (
	"github.com/chunnior/api-gateway/internal/domain"
	"github.com/gofiber/fiber/v2"
	//"github.com/chunnior/api-gateway/internal/usecase"
)

type UserHandler struct {
	userService domain.UserService
	authService domain.AuthService
}

func NewUserHandler(userService domain.UserService, authService domain.AuthService) *UserHandler {
	return &UserHandler{
		userService: userService,
		authService: authService,
	}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	baseURL := c.BaseURL()
	callbackURL := baseURL + "/api/users/callback"
	// Extrae el provider del cuerpo de la solicitud
	var body domain.LoginRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if body.Provider == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Provider is required",
		})
	}
	requestBody := domain.LoginUserServiceRequest{
		Provider:    body.Provider,
		CallbackURL: callbackURL,
	}
	loginResponse, err := h.userService.Login(requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(loginResponse)

}

func (h *UserHandler) Callback(c *fiber.Ctx) error {
	var params domain.LoginCallbackParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse query parameters",
		})
	}
	if params.Provider == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Provider is required",
		})
	}
	//	TODO: Validar demas parametros segun el proveedor
	if params.State == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "State is required",
		})
	}
	if params.Code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Code is required",
		})
	}
	callbackResponse, err := h.userService.Callback(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(callbackResponse)
}
