package events

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type CreateRoomEventHandler struct {
	log         logger.ILogger
	ctx         context.Context
	redisClient *redis.Client
}

func NewCreateRoomEventHandler(log logger.ILogger, ctx context.Context, redisClient *redis.Client) *CreateRoomEventHandler {
	return &CreateRoomEventHandler{
		log:         log,
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (h *CreateRoomEventHandler) Handle(ctx context.Context, event *CreateRoomEvent) error {
	roomId := event.RoomId
	playerId := event.PlayerId

	redisEvent := map[string]string{
		"room_id":   roomId,
		"player_id": playerId,
		"event":     "room:joined",
	}
	data, err := json.Marshal(redisEvent)
	if err != nil {
		return err
	}

	if err := h.redisClient.Publish(h.ctx, "room_events", data).Err(); err != nil {
		return err
	}

	return nil
}
