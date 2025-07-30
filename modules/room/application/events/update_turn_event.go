package events

import (
	"context"

	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type UpdateTurnEvent struct {
	RoomId string `json:"room_id"`
	Turn   int    `json:"turn"`
}

type UpdateTurnEventHandler struct {
	log logger.ILogger
	ctx context.Context
	db  *infrastructure.RoomDbContext
}

func NewUpdateTurnEventHandler(log logger.ILogger, ctx context.Context, db *infrastructure.RoomDbContext) *UpdateTurnEventHandler {
	return &UpdateTurnEventHandler{
		log: log,
		ctx: ctx,
		db:  db,
	}
}

func (h *UpdateTurnEventHandler) Handle(ctx context.Context, event *UpdateTurnEvent) error {
	go func() {
		err := h.db.GetModelDB(&domain.Room{}).
			Where("id = ?", event.RoomId).
			Update("turn", event.Turn).Error
		if err != nil {
			h.log.Error("Failed to update room turn", "error", err, "room_id", event.RoomId, "turn", event.Turn)
			return
		}
		h.log.Info("Room turn updated successfully", "room_id", event.RoomId, "turn", event.Turn)
	}()
	return nil
}
