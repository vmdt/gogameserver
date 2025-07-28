package controller

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/modules/identity/api/handler"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func IdentityRoute(echo *echo.Echo, ctx context.Context, log logger.ILogger, validator *validator.Validate) {
	group := echo.Group("/api/v1/identity")

	group.POST("/register", handler.Register(validator, ctx))
	group.POST("/login", handler.Login(validator, ctx))
}
