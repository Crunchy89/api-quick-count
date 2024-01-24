package http

import (
	"github.com/crunchy89/api-quick-count/app/domain/payloads"
	"github.com/crunchy89/api-quick-count/app/service"
	"github.com/crunchy89/api-quick-count/utils/response"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(
	userService service.UserService,
) AuthHandler {
	return AuthHandler{
		userService: userService,
	}
}

func (b *AuthHandler) Login(c echo.Context) error {
	var payload payloads.Auth
	if err := c.Bind(&payload); err != nil {
		return err
	}
	if err := c.Validate(&payload); err != nil {
		return err
	}
	data, err := b.userService.AuthService(payload)
	return response.Auto(c, data, err)
}
