package events

import (
	"context"

	boardgame_events "github.com/vmdt/gogameserver/modules/boardgame/application/events"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type EndgameEvent struct {
	RoomId   string `json:"room_id"`
	PlayerId string `json:"player_id"`
	Who      int    `json:"who"`
}

type EndgameEventHandler struct {
	log       logger.ILogger
	ctx       context.Context
	dbContext *infrastructure.RoomDbContext
}

func NewEndgameEventHandler(log logger.ILogger, ctx context.Context, dbContext *infrastructure.RoomDbContext) *EndgameEventHandler {
	return &EndgameEventHandler{
		log:       log,
		ctx:       ctx,
		dbContext: dbContext,
	}
}

func (h *EndgameEventHandler) Handle(ctx context.Context, event *boardgame_events.EndgameEvent) error {
	if err := h.dbContext.GetModelDB(&domain.Room{}).Where("id = ?", event.RoomId).Updates(map[string]interface{}{
		"is_ended": true,
		"who_win":  event.Who,
	}).Error; err != nil {
		h.log.Error("Failed to update room endgame status", "error", err, "room_id", event.RoomId)
		return err
	}
	h.log.Info("Room endgame status updated successfully", "room_id", event.RoomId, "who_win", event.Who)
	return nil
}
