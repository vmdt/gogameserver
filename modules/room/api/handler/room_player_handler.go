package handler

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	player_room_cmd "github.com/vmdt/gogameserver/modules/room/application/commands/player_room"
	"github.com/vmdt/gogameserver/modules/room/application/dtos"
)

// UpdateRoomPlayerHandler godoc
// @Summary      Update a player in a room
// @Description  Partially updates the information of a player in a room.
// @Tags         Room.Players
// @Accept       json
// @Produce      json
// @Param        room_id      path      string                                       true   "Room ID (UUID)"
// @Param        player_id    path      string                                       true   "Player ID (UUID)"
// @Param        body         body      player_room_cmd.UpdateRoomPlayerCommand     true   "Fields to update"
// @Success      200          {object}  dtos.RoomPlayerDTO
// @Failure      400          {object}  map[string]string
// @Failure      500          {object}  map[string]string
// @Router       /api/v1/room/{room_id}/players/{player_id} [put]
func UpdateRoomPlayerHandler(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req player_room_cmd.UpdateRoomPlayerCommand
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid request body"})
		}

		result, err := mediatr.Send[*player_room_cmd.UpdateRoomPlayerCommand, *dtos.RoomPlayerDTO](c.Request().Context(), &req)

		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, result)
	}
}

// KickPlayerRoomHandler handles the request to kick a player from a room.
//
// @Summary      Kick a player from a room
// @Description  Removes a player from the specified room.
// @Tags         Room.Players
// @Accept       json
// @Produce      json
// @Param        room_id      path      string                                       true   "Room ID (UUID)"
// @Param        player_id    path      string                                       true   "Player ID (UUID)"
// @Success      200   {object}  bool    "success"
// @Failure      400   {object}  map[string]string  "bad request"
// @Failure      500   {object}  map[string]string  "internal server error"
// @Router       /api/v1/room/{room_id}/players/{player_id} [delete]
func KickPlayerRoomHandler(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req player_room_cmd.KickPlayerRoomCommand
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, err.Error())
		}

		result, err := mediatr.Send[*player_room_cmd.KickPlayerRoomCommand, bool](c.Request().Context(), &req)

		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, result)
	}
}
