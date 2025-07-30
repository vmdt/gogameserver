package events

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type UpdateRoomStatusEvent struct {
	RoomId string `json:"room_id"`
	Status string `json:"status"`
}

type UpdateRoomStatusEventHandler struct {
	log         logger.ILogger
	ctx         context.Context
	redisClient *redis.Client
}

func NewUpdateRoomStatusEventHandler(log logger.ILogger, ctx context.Context, redisClient *redis.Client) *UpdateRoomStatusEventHandler {
	return &UpdateRoomStatusEventHandler{
		log:         log,
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (h *UpdateRoomStatusEventHandler) Handle(ctx context.Context, event *UpdateRoomStatusEvent) error {
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
