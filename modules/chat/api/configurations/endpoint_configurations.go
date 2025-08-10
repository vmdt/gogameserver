package configurations

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/modules/chat/api/controller"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigEndpoints(
	log logger.ILogger,
	ctx context.Context,
	echo *echo.Echo,
	validator *validator.Validate,
) error {
	controller.ChatRoutes(echo, ctx, log, validator)
	log.Info("Chat routes configured successfully")
	return nil
}
