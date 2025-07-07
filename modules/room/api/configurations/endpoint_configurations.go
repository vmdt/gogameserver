package configurations

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/modules/room/api/controller"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigEndpoints(log logger.ILogger, echo *echo.Echo, ctx context.Context) error {
	controller.RoomRoute(echo, ctx, log)
	return nil
}
