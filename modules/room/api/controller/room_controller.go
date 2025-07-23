package controller

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/modules/room/api/handler"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func RoomRoute(echo *echo.Echo, ctx context.Context, log logger.ILogger, validator *validator.Validate) {
	group := echo.Group("/api/v1/room")

	group.POST("/player/create", handler.PlayerCreateRoomHandler(validator, ctx))
	group.POST("/player/join", handler.PlayerJoinRoomHandler(validator, ctx))

	group.PUT("/status", handler.UpdateRoomStatusHandler(validator, ctx))
	group.POST("/create", handler.CreateRoomHandler())
	group.GET("/:id", handler.GetRoomHandler())
}
