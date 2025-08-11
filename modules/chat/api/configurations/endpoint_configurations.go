package configurations

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/modules/chat/api/controller"
	"github.com/vmdt/gogameserver/pkg/auth"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigEndpoints(
	log logger.ILogger,
	ctx context.Context,
	echo *echo.Echo,
	validator *validator.Validate,
	jwtService auth.IJwtService,
) error {
	controller.ChatRoutes(echo, ctx, log, validator, jwtService)
	log.Info("Chat routes configured successfully")
	return nil
}
