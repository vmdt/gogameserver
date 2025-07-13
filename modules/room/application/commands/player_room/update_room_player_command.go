package player_room_cmd

import (
	"context"
	"time"

	"github.com/vmdt/gogameserver/modules/room/application/dtos"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type UpdateRoomPlayerCommand struct {
	RoomId         string     `param:"room_id" json:"-" validate:"required"`
	PlayerId       string     `param:"player_id" json:"-" validate:"required"`
	IsReady        *bool      `json:"is_ready,omitempty"`
	IsDisconnected *bool      `json:"is_disconnected,omitempty"`
	DisconnectedAt *time.Time `json:"disconnected_at,omitempty"`
	IsHost         *bool      `json:"is_host,omitempty"`
}

func NewUpdateRoomPlayerCommand() *UpdateRoomPlayerCommand {
	return &UpdateRoomPlayerCommand{}
}

type UpdateRoomPlayerCommandHandler struct {
	log logger.ILogger
	ctx context.Context
	db  *infrastructure.RoomDbContext
}

func NewUpdateRoomPlayerCommandHandler(log logger.ILogger, ctx context.Context, db *infrastructure.RoomDbContext) *UpdateRoomPlayerCommandHandler {
	return &UpdateRoomPlayerCommandHandler{
		log: log,
		ctx: ctx,
		db:  db,
	}
}

func (h *UpdateRoomPlayerCommandHandler) Handle(ctx context.Context, command *UpdateRoomPlayerCommand) (*dtos.RoomPlayerDTO, error) {
	roomId := command.RoomId
	playerId := command.PlayerId
	var existing domain.RoomPlayer
	if err := h.db.GetModelDB(&domain.RoomPlayer{}).
		Where("room_id = ? AND player_id = ?", roomId, playerId).
		Preload("Player").
		First(&existing).Error; err != nil {
		h.log.Error("RoomPlayer not found", "error", err)
		return nil, err
	}

	updates := map[string]interface{}{}

	if command.IsReady != nil {
		updates["is_ready"] = *command.IsReady
	}
	if command.IsDisconnected != nil {
		updates["is_disconnected"] = *command.IsDisconnected
	}
	if command.DisconnectedAt != nil {
		updates["disconnected_at"] = *command.DisconnectedAt
	}
	if command.IsHost != nil {
		updates["is_host"] = *command.IsHost
	}

	if len(updates) == 0 {
		h.log.Info("No fields to update for RoomPlayer", "room_id", roomId, "player_id", playerId)
		return &dtos.RoomPlayerDTO{
			RoomId:         existing.RoomId.String(),
			PlayerId:       existing.PlayerId.String(),
			IsReady:        existing.IsReady,
			IsDisconnected: existing.IsDisconnected,
			DisconnectedAt: existing.DisconnectedAt,
			IsHost:         existing.IsHost,
			Player:         existing.Player.ToDTO(),
		}, nil
	}

	var updated domain.RoomPlayer
	if err := h.db.GetModelDB(&domain.RoomPlayer{}).
		Where("room_id = ? AND player_id = ?", roomId, playerId).
		Updates(updates).
		Preload("Player").
		First(&updated).Error; err != nil {
		h.log.Error("Failed to update RoomPlayer", "error", err)
		return nil, err
	}
	h.log.Info("RoomPlayer updated successfully", "room_id", roomId, "player_id", playerId)
	return &dtos.RoomPlayerDTO{
		RoomId:         updated.RoomId.String(),
		PlayerId:       updated.PlayerId.String(),
		IsReady:        updated.IsReady,
		IsDisconnected: updated.IsDisconnected,
		DisconnectedAt: updated.DisconnectedAt,
		IsHost:         updated.IsHost,
		Player:         updated.Player.ToDTO(),
	}, nil
}
