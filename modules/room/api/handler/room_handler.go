package handler

import (
	"context"
	"errors"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/room/application/commands"
	"github.com/vmdt/gogameserver/modules/room/application/dtos"
	"github.com/vmdt/gogameserver/modules/room/application/query"
)

func CreateRoomHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := &commands.CreateRoomCommand{}
		if err := c.Bind(request); err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid request body"})
		}

		cmd := commands.NewCreateRoomCommand(request.Status)

		result, err := mediatr.Send[*commands.CreateRoomCommand, *dtos.RoomDTO](c.Request().Context(), cmd)

		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return c.JSON(200, result)
	}
}

func GetRoomHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		roomID := c.Param("id")
		q := query.NewGetRoomQuery(roomID)

		result, err := mediatr.Send[*query.GetRoomQuery, *dtos.RoomInformationDTO](c.Request().Context(), q)

		if err != nil {
			return c.JSON(400, map[string]string{"error": errors.Unwrap(err).Error()})
		}

		return c.JSON(200, result)
	}
}

// @Summary      Player creates a new room
// @Description  Allows a player to create a new room with a name and user ID
// @Tags         Room.Player
// @Accept       json
// @Produce      json
// @Param        request  body      commands.PlayerCreateRoomCommand  true  "Player Create Room Request"
// @Success      200      {object}  dtos.RoomPlayerDTO
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /api/v1/room/player/create [post]
func PlayerCreateRoomHandler(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := &commands.PlayerCreateRoomCommand{}
		if err := c.Bind(request); err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid request body"})
		}

		if err := validator.StructCtx(ctx, request); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}

		result, err := mediatr.Send[*commands.PlayerCreateRoomCommand, *dtos.RoomPlayerDTO](ctx, request)

		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return c.JSON(200, result)
	}
}

// @Summary      Player joins a room
// @Description  Allows a player to join an existing room with a name, user ID,
// @Tags         Room.Player
// @Accept       json
// @Produce      json
// @Param        request  body      commands.JoinRoomCommand  true  "Player Join Room"
// @Success      200      {object}  dtos.RoomPlayerDTO
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /api/v1/room/player/join [post]
func PlayerJoinRoomHandler(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := &commands.JoinRoomCommand{}
		if err := c.Bind(request); err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid request body"})
		}

		if err := validator.StructCtx(ctx, request); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}

		cmd := commands.NewJoinRoomCommand(request.Name, request.UserId, request.PlayerId, request.RoomId)

		result, err := mediatr.Send[*commands.JoinRoomCommand, *dtos.RoomPlayerDTO](ctx, cmd)

		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return c.JSON(200, result)
	}
}

// @Summary      Update room status
// @Description  Allows updating the status of a room by its ID
// @Tags         Room
// @Accept       json
// @Produce      json
// @Param        request  body      commands.UpdateRoomStatusCommand  true  "Update Room Status Request"
// @Success      200      {object}  dtos.RoomDTO
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /api/v1/room/status [put]
func UpdateRoomStatusHandler(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := &commands.UpdateRoomStatusCommand{}
		if err := c.Bind(request); err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid request body"})
		}

		if err := validator.StructCtx(ctx, request); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}

		cmd := commands.NewUpdateRoomStatusCommand(request.RoomId, request.Status)

		result, err := mediatr.Send[*commands.UpdateRoomStatusCommand, *dtos.RoomDTO](ctx, cmd)

		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return c.JSON(200, result)
	}
}
