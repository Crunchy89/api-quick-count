package middleware

import (
	"fmt"
	"net/http"
	"strings"

	_jwt "github.com/crunchy89/api-quick-count/utils/jwt"
	"github.com/crunchy89/api-quick-count/utils/response"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type InitMiddleware struct {
	_jwt _jwt.JWTHelper
}

func NewInitMiddleware(jwt _jwt.JWTHelper) InitMiddleware {
	return InitMiddleware{_jwt: jwt}
}

func (b *InitMiddleware) VerifyAccess() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token == "" && !strings.HasPrefix(token, "Bearer") {
				return response.ErrorWithMessageStatus(c, http.StatusUnauthorized, "authorization not valid")

			}

			splintedToken := strings.Split(token, " ")
			if len(splintedToken) != 2 {
				return response.ErrorWithMessageStatus(c, http.StatusUnauthorized, "authorization not valid")
			}

			jwtClaims, err := b._jwt.ValidateToken(splintedToken[1])
			if err != nil {
				return response.ErrorWithMessageStatus(c, http.StatusUnauthorized, "authorization not valid")
			}

			customClaims, ok := jwtClaims.Claims.(jwt.MapClaims)
			if !ok {
				return response.ErrorWithMessageStatus(c, http.StatusUnauthorized, "authorization not valid")
			}
			c.Set("uuid", fmt.Sprintf("%v", customClaims["uuid"]))

			// this we accept and continue the request
			return next(c)
		}
	}
}
