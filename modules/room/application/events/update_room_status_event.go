package events

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type UpdateRoomStatusEvent struct {
	RoomId string `json:"room_id"`
	Status string `json:"status"`
}

type UpdateRoomStatusEventHandler struct {
	log         logger.ILogger
	ctx         context.Context
	db          *infrastructure.RoomDbContext
	redisClient *redis.Client
}

func NewUpdateRoomStatusEventHandler(log logger.ILogger, ctx context.Context, db *infrastructure.RoomDbContext, redisClient *redis.Client) *UpdateRoomStatusEventHandler {
	return &UpdateRoomStatusEventHandler{
		log:         log,
		ctx:         ctx,
		db:          db,
		redisClient: redisClient,
	}
}

func (h *UpdateRoomStatusEventHandler) Handle(ctx context.Context, event *UpdateRoomStatusEvent) error {
	go func(db *infrastructure.RoomDbContext, event *UpdateRoomStatusEvent) {
		if event.Status == "setup" {
			if err := db.GetModelDB(&domain.BattleshipOptions{}).Where("room_id = ?", event.RoomId).Update("start_place_at", time.Now()).Error; err != nil {
				h.log.Error("Failed to update room status to setup", "error", err, "room_id", event.RoomId)
				return
			}
			h.log.Info("Room status updated to setup", "room_id", event.RoomId)
		}
	}(h.db, event)

	redisEvent := map[string]string{
		"room_id": event.RoomId,
		"status":  event.Status,
		"event":   "room:started",
	}
	data, err := json.Marshal(redisEvent)
	if err != nil {
		h.log.Error("Failed to marshal redis event", "error", err)
		return err
	}
	if err := h.redisClient.Publish(h.ctx, "room_events", data).Err(); err != nil {
		h.log.Error("Failed to publish redis event", "error", err)
		return err
	}
	return nil
}
