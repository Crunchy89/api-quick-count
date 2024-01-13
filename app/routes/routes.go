package routes

import (
	"github.com/crunchy89/api-quick-count/app/http"
	"github.com/crunchy89/api-quick-count/app/middleware"
	"github.com/labstack/echo/v4"
)

type V1Routes struct {
	echo        *echo.Echo
	middleware  middleware.InitMiddleware
	roleHandler http.RoleHandler
}

func NewV1Routes(
	echo *echo.Echo,
	middleware middleware.InitMiddleware,
	roleHandler http.RoleHandler,
) V1Routes {
	return V1Routes{
		echo:        echo,
		middleware:  middleware,
		roleHandler: roleHandler,
	}
}

func (r *V1Routes) Routes() {
	api := r.echo.Group("api").Group("/v1")
	roleApi := api.Group("/role")
	roleApi.Use(r.middleware.VerifyAccess())
	roleApi.GET("fetch/all", r.roleHandler.GetRole)

}
