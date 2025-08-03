package controller

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/modules/room/api/handler"
	"github.com/vmdt/gogameserver/pkg/auth"
	"github.com/vmdt/gogameserver/pkg/logger"
	"github.com/vmdt/gogameserver/server/middlewares"
)

func RoomRoute(echo *echo.Echo, ctx context.Context, log logger.ILogger, validator *validator.Validate, jwtService auth.IJwtService) {
	group := echo.Group("/api/v1/room")

	group.POST("/player/create", handler.PlayerCreateRoomHandler(validator, ctx), middlewares.JwtAuthVerifier(jwtService, ctx))
	group.POST("/player/join", handler.PlayerJoinRoomHandler(validator, ctx), middlewares.JwtAuthVerifier(jwtService, ctx))

	group.PUT("/:room_id/set-who-win", handler.SetWhoWinHandler(validator, ctx), middlewares.JwtAuthVerifier(jwtService, ctx))

	// Battleship options
	group.PUT("/:room_id/battleship-options", handler.UpdateBattleshipOptions(validator, ctx), middlewares.JwtAuthVerifier(jwtService, ctx))

	group.PUT("/status", handler.UpdateRoomStatusHandler(validator, ctx))
	group.POST("/create", handler.CreateRoomHandler())
	group.GET("/:id", handler.GetRoomHandler(), middlewares.JwtAuthVerifier(jwtService, ctx))
}
