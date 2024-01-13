package http

import (
	"github.com/crunchy89/api-quick-count/app/domain/payloads"
	"github.com/crunchy89/api-quick-count/app/service"
	"github.com/crunchy89/api-quick-count/utils/response"
	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(
	roleService service.RoleService,
) RoleHandler {
	return RoleHandler{
		roleService: roleService,
	}
}

func (b *RoleHandler) GetRole(c echo.Context) error {
	var payload payloads.CreateRolePayload
	if err := c.Bind(&payload); err != nil {
		return err
	}
	if err := c.Validate(&payload); err != nil {
		return err
	}
	data, err := b.roleService.GetAllRole()
	return response.Auto(c, data, err)
}
