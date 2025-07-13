package controller

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/modules/room/api/handler"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func RoomPlayerRoute(echo *echo.Echo, ctx context.Context, log logger.ILogger, validator *validator.Validate) {
	group := echo.Group("/api/v1/room")
	group.PUT("/:room_id/players/:player_id", handler.UpdateRoomPlayerHandler(validator, ctx))
}
