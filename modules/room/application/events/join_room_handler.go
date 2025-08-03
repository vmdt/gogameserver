package events

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type JoinRoomEventHandler struct {
	log         logger.ILogger
	ctx         context.Context
	redisClient *redis.Client
}

func NewJoinRoomEventHandler(log logger.ILogger, ctx context.Context, redisClient *redis.Client) *JoinRoomEventHandler {
	return &JoinRoomEventHandler{
		log:         log,
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (h *JoinRoomEventHandler) Handle(ctx context.Context, event *JoinRoomEvent) error {
	roomId := event.RoomId
	playerId := event.PlayerId

	redisEvent := map[string]string{
		"room_id":   roomId,
		"player_id": playerId,
		"user_id":   event.UserId, // Include user_id in the event
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
