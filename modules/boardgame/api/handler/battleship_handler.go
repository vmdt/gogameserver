package handler

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/boardgame/application/commands"
	"github.com/vmdt/gogameserver/modules/boardgame/application/dtos"
	"github.com/vmdt/gogameserver/modules/boardgame/application/queries"
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
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	}
}

// GetBattleShipBoardHandler handles the retrieval of a Battleship game board.
// @Summary      Get Battleship Board
// @Description  Retrieves the Battleship game board for a specific player and room.
// @Tags         Board.Battleship
// @Accept       json
// @Produce      json
// @Param        player_id  path      string  true  "Player ID"
// @Param        room_id    path      string  true  "Room ID"
// @Success      200  {object}   dtos.BattleshipGame
// @Failure      400  {string}   string  "Invalid request or Validation error
// @Failure      500  {string}   string  "Internal server error"
// @Router       /api/v1/boardgame/battleship/room/{room_id}/player/{player_id} [get]
func GetBattleShipBoardHandler(
	validator *validator.Validate,
	log logger.ILogger,
	ctx context.Context,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query queries.GetBattleshipBoardQuery
		if err := c.Bind(&query); err != nil {
			log.Error("Failed to bind request", "error", err)
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}

		if err := validator.Struct(query); err != nil {
			log.Error("Validation failed", "error", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		cmd := queries.NewGetBattleshipBoardQuery(
			query.PlayerId,
			query.RoomId,
		)
		result, err := mediatr.Send[*queries.GetBattleshipBoardQuery, *dtos.BattleshipGame](ctx, cmd)
		if err != nil {
			log.Error("Failed to get battleship board", "error", err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	}
}

// AttackBattleShipHandler handles the attack on a Battleship game board.
// @Summary      Attack Battleship Board
// @Description  Attacks a position on the Battleship game board for a specific player and
// @Tags         Board.Battleship
// @Accept       json
// @Produce      json
// @Param        AttackBattleShipCommand  body      commands.AttackBattleShipCommand  true  "Attack Battleship Command Data"
// @Success      200  {object}   bool
// @Failure      400  {string}   string  "Invalid request or Validation error"
// @Failure      500  {string}   string  "Internal server error"
// @Router       /api/v1/boardgame/battleship/attack [put]
func AttackBattleShipHandler(
	validator *validator.Validate,
	log logger.ILogger,
	ctx context.Context,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request commands.AttackBattleShipCommand
		if err := c.Bind(&request); err != nil {
			log.Error("Failed to bind request", "error", err)
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}
		if err := validator.Struct(request); err != nil {
			log.Error("Validation failed", "error", err)
			return c.JSON(http.StatusBadRequest, "Validation error")
		}

		cmd := commands.NewAttackBattleShipCommand(
			request.PlayerId,
			request.RoomId,
			request.Position,
		)
		result, err := mediatr.Send[*commands.AttackBattleShipCommand, bool](ctx, cmd)
		if err != nil {
			log.Error("Failed to attack battleship", "error", err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, result)
	}
}

// CheckWhoWinHandler handles the check for who wins the Battleship game.
// @Summary      Check Who Wins
// @Description  Checks who wins the Battleship game for a specific room.
// @Tags         Board.Battleship
// @Accept       json
// @Produce      json
// @Param        room_id  path      string  true  "Room ID"
// @Success      200  {object}   dtos.WhoWinDTO
// @Failure      400  {string}   string  "Invalid request or Validation error"
// @Failure      404  {string}   string  "Room not found"
// @Failure      500  {string}   string  "Internal server error"
// @Router       /api/v1/boardgame/battleship/room/{room_id}/check-who-win [get]
func CheckWhoWin(
	validator *validator.Validate,
	log logger.ILogger,
	ctx context.Context,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query queries.CheckWhoWinQuery
		if err := c.Bind(&query); err != nil {
			log.Error("Failed to bind request", "error", err)
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}

		if err := validator.Struct(query); err != nil {
			log.Error("Validation failed", "error", err)
			return c.JSON(http.StatusBadRequest, "Validation error")
		}

		cmd := queries.NewCheckWhoWinQuery(query.RoomId)
		result, err := mediatr.Send[*queries.CheckWhoWinQuery, *dtos.WhoWinDTO](ctx, cmd)
		if err != nil {
			log.Error("Failed to check who win", "error", err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	}
}

// CheckSunkShipStatusHandler handles the checking of sunk ship status.
// @Summary      Check Sunk Ship Status
// @Description  Checks the status of sunk ships for a specific player in a room.
// @Tags         Board.Battleship
// @Accept       json
// @Produce      json
// @Param        CheckSunkShipStatusQuery  body      queries.CheckSunkShipStatusQuery  true  "Check Sunk Ship Status Query Data"
// @Success      200  {object}   dtos.SunkShipsDTO
// @Failure      400  {string}   string  "Invalid request or Validation error"
// @Failure      404  {string}   string  "Player or room not found"
// @Failure      500  {string}   string  "Internal server error"
// @Router       /api/v1/boardgame/battleship/room/{room_id}/player/{player_id}/check-sunk-ships [get]
func CheckSunkShipStatus(
	validator *validator.Validate,
	log logger.ILogger,
	ctx context.Context,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query queries.CheckSunkShipStatusQuery
		if err := c.Bind(&query); err != nil {
			log.Error("Failed to bind request", "error", err)
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}

		if err := validator.Struct(query); err != nil {
			log.Error("Validation failed", "error", err)
			return c.JSON(http.StatusBadRequest, "Validation error")
		}

		cmd := queries.NewCheckSunkShipStatusQuery(query.RoomId, query.PlayerId)
		result, err := mediatr.Send[*queries.CheckSunkShipStatusQuery, *dtos.SunkShipsDTO](ctx, cmd)
		if err != nil {
			log.Error("Failed to check sunk ship status", "error", err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	}
}
