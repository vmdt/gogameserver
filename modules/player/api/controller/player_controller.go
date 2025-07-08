package controller

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/modules/player/api/handler"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func PlayerRoute(echo *echo.Echo, ctx context.Context, log logger.ILogger, validator *validator.Validate) error {
	group := echo.Group("/api/v1/player")

	group.POST("/create", handler.CreatePlayerHandler(validator, log, ctx))

	log.Info("Player routes configured successfully")
	return nil
}
