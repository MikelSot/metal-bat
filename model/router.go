package model

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

type RouterSpecification struct {
	Config   Configuration
	Api      *echo.Echo
	Logger   Logger
	DBEngine string
	DB       *sql.DB
	//Auth     authorization.AuthMiddleware
	//Scope    authorization.ScopeMiddleware
}
