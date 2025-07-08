package handler

import (
	"context"

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

		result, err := mediatr.Send[*query.GetRoomQuery, *dtos.RoomDTO](c.Request().Context(), q)

		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
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

		cmd := commands.NewPlayerCreateRoomCommand(request.Name, request.UserId)

		result, err := mediatr.Send[*commands.PlayerCreateRoomCommand, *dtos.RoomPlayerDTO](ctx, cmd)

		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return c.JSON(200, result)
	}
}
