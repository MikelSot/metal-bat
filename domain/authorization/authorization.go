package authorization

import "github.com/labstack/echo/v4"

type AuthMiddleware interface {
	ValidateToken(next echo.HandlerFunc) echo.HandlerFunc
}
