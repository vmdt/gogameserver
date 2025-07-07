package controller

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/modules/room/api/handler"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func RoomRoute(echo *echo.Echo, ctx context.Context, log logger.ILogger) {
	group := echo.Group("/api/v1/room")

	group.POST("/create", handler.CreateRoomHandler())
	group.GET("/:id", handler.GetRoomHandler())
}
