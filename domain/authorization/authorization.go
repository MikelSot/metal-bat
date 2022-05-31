package authorization

import "github.com/labstack/echo/v4"

type AuthMiddleware interface {
	ValidateServiceAccessCode(next echo.HandlerFunc) echo.HandlerFunc
	ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc
}
