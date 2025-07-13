package player_room_cmd

import (
	"context"

	"github.com/google/uuid"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type InternalCreateRoomPlayerCommand struct {
	RoomId   string `json:"room_id"`
	PlayerId string `json:"player_id"`
	IsHost   bool   `json:"is_host"`
}

func NewInternalCreateRoomPlayerCommand(roomId, playerId string, isHost bool) *InternalCreateRoomPlayerCommand {
	return &InternalCreateRoomPlayerCommand{
		RoomId:   roomId,
		PlayerId: playerId,
		IsHost:   isHost,
	}
}

type InternalCreateRoomPlayerCommandHandler struct {
	log logger.ILogger
	ctx context.Context
	db  *infrastructure.RoomDbContext
}

func NewInternalCreateRoomPlayerCommandHandler(log logger.ILogger, ctx context.Context, db *infrastructure.RoomDbContext) *InternalCreateRoomPlayerCommandHandler {
	return &InternalCreateRoomPlayerCommandHandler{
		log: log,
		ctx: ctx,
		db:  db,
	}
}

func (h *InternalCreateRoomPlayerCommandHandler) Handle(ctx context.Context, command *InternalCreateRoomPlayerCommand) (*domain.RoomPlayer, error) {
	roomPlayer := &domain.RoomPlayer{
		RoomId:   uuid.MustParse(command.RoomId),
		PlayerId: uuid.MustParse(command.PlayerId),
		IsHost:   command.IsHost,
	}
	if err := h.db.GetModelDB(&domain.RoomPlayer{}).Create(roomPlayer).Error; err != nil {
		h.log.Error("Failed to create room player", "error", err)
		return nil, err
	}
	h.log.Info("Room player created successfully", "room_id", command.RoomId, "player_id", command.PlayerId)
	return roomPlayer, nil
}
