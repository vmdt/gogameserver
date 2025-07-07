package handler

import (
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
