package handler

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/chat/application/commands"
	"github.com/vmdt/gogameserver/modules/chat/application/dtos"
	"github.com/vmdt/gogameserver/modules/chat/application/queries"
)

// CreateRoomChatHandler handles the creation of a chat room
// @Summary      Create a new chat room
// @Description  Creates a new chat room and returns the chat details.
// @Tags         Chat
// @Accept       json
// @Produce      json
// @Param        body         body      commands.CreateChatCommand     true   "Chat room details"
// @Success      201          {object}  dtos.ChatDTO
// @Failure      400          {object}  map[string]string
// @Failure      500          {object}  map[string]string
// @Router       /api/v1/chat/rooms [post]
func CreateRoomChat(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req commands.CreateChatCommand
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}
		if err := validator.Struct(req); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}

		result, err := mediatr.Send[*commands.CreateChatCommand, *dtos.ChatDTO](ctx, &req)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return c.JSON(201, result)
	}
}

// GetRoomChatHandler retrieves a chat room by its ID
// @Summary      Get chat room by ID
// @Description  Retrieves a chat room by its ID and returns the chat details.
// @Tags         Chat
// @Accept       json
// @Produce      json
// @Param        room_id       path      string                        true   "Chat room ID"
// @Success      200          {object}  dtos.ChatDTO
// @Failure      400          {object}  map[string]string
// @Failure      404          {object}  map[string]string
// @Failure      500          {object}  map[string]string
// @Router       /api/v1/chat/rooms/{room_id} [get]
func GetRoomChat(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req queries.GetChatQuery
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}
		if err := validator.Struct(req); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}

		result, err := mediatr.Send[*queries.GetChatQuery, *dtos.ChatDTO](ctx, &req)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return c.JSON(200, result)
	}
}

// SendChatMessage handles sending a chat message
// @Summary      Send a chat message
// @Description  Sends a chat message to a specific chat room and returns the updated chat details
// @Tags         Chat
// @Accept       json
// @Produce      json
// @Param        room_id      path      string                        true   "Chat room ID"
// @Param        body         body      commands.ChatMessageCommand     true   "Chat message details"
// @Success      200          {object}  dtos.ChatDTO
// @Failure      400          {object}  map[string]string
// @Failure      404          {object}  map[string]string
// @Failure      500          {object}  map[string]string
// @Router       /api/v1/chat/rooms/{room_id}/messages [post]
func SendChatMessage(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req commands.ChatMessageCommand
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}
		if err := validator.Struct(req); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}

		result, err := mediatr.Send[*commands.ChatMessageCommand, *dtos.ChatDTO](ctx, &req)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return c.JSON(200, result)
	}
}
