//go:build wireinject
// +build wireinject

package main

import (
	"github.com/centrifugal/centrifuge"
	"github.com/crunchy89/api-quick-count/app/http"
	"github.com/crunchy89/api-quick-count/app/middleware"
	"github.com/crunchy89/api-quick-count/app/repository"
	"github.com/crunchy89/api-quick-count/app/routes"
	"github.com/crunchy89/api-quick-count/app/service"
	"github.com/crunchy89/api-quick-count/utils/jwt"
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
		repository.NewUserRepository,
		service.NewRoleService,
		service.NewUserService,
		http.NewRoleHandler,
		http.NewAuthHandler,
		routes.NewV1Routes,
	)
	return routes.V1Routes{}
}
