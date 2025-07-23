package controller

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/modules/boardgame/api/handler"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func BattleshipRoute(
	echo *echo.Echo,
	ctx context.Context,
	log logger.ILogger,
	validator *validator.Validate,
) error {
	group := echo.Group("/api/v1/boardgame/battleship")

	group.GET("/room/:room_id/player/:player_id", handler.GetBattleShipBoardHandler(validator, log, ctx))
	group.POST("/", handler.CreateBattleShipBoardHandler(validator, log, ctx))
	return nil
}
