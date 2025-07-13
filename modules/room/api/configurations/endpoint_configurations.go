package configurations

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/modules/room/api/controller"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigEndpoints(log logger.ILogger, echo *echo.Echo, ctx context.Context, validator *validator.Validate) error {
	controller.RoomRoute(echo, ctx, log, validator)
	controller.RoomPlayerRoute(echo, ctx, log, validator)
	return nil
}
