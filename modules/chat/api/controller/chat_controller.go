package controller

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/modules/chat/api/handler"
	"github.com/vmdt/gogameserver/pkg/auth"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ChatRoutes(echo *echo.Echo, ctx context.Context, log logger.ILogger, validator *validator.Validate, jwtService auth.IJwtService) {
	group := echo.Group("/api/v1/chat")

	group.POST("/rooms", handler.CreateRoomChat(validator, ctx))
	group.POST("/rooms/:room_id/messages", handler.SendChatMessage(validator, ctx))
	group.GET("/rooms/:room_id", handler.GetRoomChat(validator, ctx))
}
