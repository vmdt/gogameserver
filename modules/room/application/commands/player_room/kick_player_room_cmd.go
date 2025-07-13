package player_room_cmd

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	player_commands "github.com/vmdt/gogameserver/modules/player/application/commands"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type KickPlayerRoomCommand struct {
	RoomId   string `param:"room_id" validate:"required" json:"-"`
	PlayerId string `param:"player_id" validate:"required" json:"-"`
}

func NewKickPlayerRoomCommand(roomId, playerId string) *KickPlayerRoomCommand {
	return &KickPlayerRoomCommand{
		RoomId:   roomId,
		PlayerId: playerId,
	}
}

type KickPlayerRoomCommandHandler struct {
	log logger.ILogger
	ctx context.Context
	db  *infrastructure.RoomDbContext
}

func NewKickPlayerRoomCommandHandler(log logger.ILogger, ctx context.Context, db *infrastructure.RoomDbContext) *KickPlayerRoomCommandHandler {
	return &KickPlayerRoomCommandHandler{
		log: log,
		ctx: ctx,
		db:  db,
	}
}

func (h *KickPlayerRoomCommandHandler) Handle(ctx context.Context, command *KickPlayerRoomCommand) (bool, error) {
	result, err := mediatr.Send[*player_commands.DeletePlayerCommand, bool](ctx, player_commands.NewDeletePlayerCommand(command.PlayerId))
	if err != nil {
		h.log.Error("Failed to kick player from room", "error", err)
		return false, err
	}
	return result, nil
}
