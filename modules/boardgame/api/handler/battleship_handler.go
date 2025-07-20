package handler

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/boardgame/application/commands"
	"github.com/vmdt/gogameserver/modules/boardgame/application/dtos"
	"github.com/vmdt/gogameserver/pkg/logger"
)

// CreateBattleShipBoardHandler handles the creation of a Battleship game board.
//
// @Summary      Create Battleship Board
// @Description  Creates a new Battleship game board with the provided player, room, ships, and shots information.
// @Tags         Board.Battleship
// @Accept       json
// @Produce      json
// @Param        BattleshipGame  body      dtos.BattleshipGame  true  "Battleship Game Data"
// @Success      200  {object}   dtos.BattleshipGame
// @Failure      400  {string}   string  "Invalid request or Validation error"
// @Failure      500  {string}   string  "Internal server error"
// @Router       /api/v1/boardgame/battleship [post]
func CreateBattleShipBoardHandler(
	validator *validator.Validate,
	log logger.ILogger,
	ctx context.Context,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		var command dtos.BattleshipGame
		if err := c.Bind(&command); err != nil {
			log.Error("Failed to bind request", "error", err)
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}

		if err := validator.Struct(command); err != nil {
			log.Error("Validation failed", "error", err)
			return c.JSON(http.StatusBadRequest, "Validation error")
		}

		cmd := commands.NewCreateBattleShipBoardCommand(
			command.PlayerId,
			command.RoomId,
			command.Ships,
			command.Shots,
		)
		result, err := mediatr.Send[*commands.CreateBattleShipBoardCommand, *dtos.BattleshipGame](ctx, cmd)
		if err != nil {
			log.Error("Failed to create battleship board", "error", err)
			return c.JSON(http.StatusInternalServerError, "Internal server error")
		}

		return c.JSON(http.StatusOK, result)
	}
}
