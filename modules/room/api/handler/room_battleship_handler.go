package handler

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	battleship_options_cmd "github.com/vmdt/gogameserver/modules/room/application/commands/battleship_options"
	"github.com/vmdt/gogameserver/modules/room/application/dtos"
)

// UpdateBattleshipOptions updates the battleship options for a room
// @Summary      Update Battleship Options
// @Description  Allows updating the battleship options for a specific room
// @Tags         Room.Battleship
// @Accept       json
// @Produce      json
// @Param        request  body      battleship_options_cmd.UpdateBattleshipOptionsCmd true "Update Battleship Options Request"
// @Success      200      {object}  dtos.BattleshipOptionsDTO
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /api/v1/room/{room_id}/battleship-options [put]
func UpdateBattleshipOptions(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cmd battleship_options_cmd.UpdateBattleshipOptionsCmd
		if err := c.Bind(&cmd); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		if err := validator.Struct(cmd); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		result, err := mediatr.Send[*battleship_options_cmd.UpdateBattleshipOptionsCmd, *dtos.BattleshipOptionsDTO](ctx, &cmd)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	}
}
