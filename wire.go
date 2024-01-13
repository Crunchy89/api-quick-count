//go:build wireinject
// +build wireinject

package main

import (
	"github.com/centrifugal/centrifuge"
	"github.com/crunchy89/api-techcode/app/v1/http"
	"github.com/crunchy89/api-techcode/app/v1/middleware"
	"github.com/crunchy89/api-techcode/app/v1/repository"
	"github.com/crunchy89/api-techcode/app/v1/routes"
	"github.com/crunchy89/api-techcode/app/v1/service"
	"github.com/crunchy89/api-techcode/utils/jwt"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitV1Routes(
	socket *centrifuge.Node,
	db *gorm.DB,
	logger *logrus.Entry,
	echo *echo.Echo,
) routes.V1Routes {
	wire.Build(
		jwt.NewJWTService,
		middleware.NewInitMiddleware,
		repository.NewRoleRepository,
		service.NewRoleService,
		http.NewRoleHandler,
		routes.NewV1Routes,
	)
	return routes.V1Routes{}
}
