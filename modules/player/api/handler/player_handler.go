package handler

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/player/application/commands"
	"github.com/vmdt/gogameserver/modules/player/application/dtos"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func CreatePlayerHandler(validator *validator.Validate, log logger.ILogger, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := &commands.CreatePlayerCommand{}
		if err := c.Bind(request); err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid request body"})
		}
		cmd := commands.NewCreatePlayerCommand(request.Name, request.UserId)

		if err := validator.StructCtx(ctx, cmd); err != nil {
			log.Error("Validation failed for CreatePlayerCommand", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		result, err := mediatr.Send[*commands.CreatePlayerCommand, *dtos.PlayerDTO](ctx, cmd)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusCreated, result)
	}
}
