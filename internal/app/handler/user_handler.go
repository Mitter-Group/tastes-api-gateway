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
	token, refresh_token, err := h.authService.GenerateTokens(callbackResponse)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"id":            callbackResponse.ID,
		"token":         token,
		"refresh_token": refresh_token,
	})
}

func (h *UserHandler) RefreshToken(c *fiber.Ctx) error {
	var body domain.RefreshTokenRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if body.RefreshToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Refresh token is required",
		})
	}
	userPayload, err := h.authService.ValidateToken(body.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid refresh token",
		})
	}
	token, refresh_token, err := h.authService.GenerateTokens(*userPayload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"token":         token,
		"refresh_token": refresh_token,
	})
}

func (h *UserHandler) HandleDataInfo(c *fiber.Ctx) error {
	provider, dataType := c.Params("provider"), c.Params("dataType")
	if provider == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Provider is required",
		})
	}
	if dataType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data type is required",
		})
	}
	userPayload := c.Locals("user").(*domain.UserPayload)
	params := domain.DataInfoParams{
		Provider: provider,
		DataType: dataType,
		UserID:   userPayload.ProviderUserID,
	}
	dataInfoResponse, err := h.userService.DataInfo(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(dataInfoResponse)
}
