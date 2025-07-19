package events

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type KickPlayerRoomEvent struct {
	RoomId   string `json:"room_id"`
	PlayerId string `json:"player_id"`
}

func NewKickPlayerRoomEvent(roomId, playerId string) *KickPlayerRoomEvent {
	return &KickPlayerRoomEvent{
		RoomId:   roomId,
		PlayerId: playerId,
	}
}

type KickPlayerRoomEventHandler struct {
	log         logger.ILogger
	ctx         context.Context
	redisClient *redis.Client
}

func NewKickPlayerRoomEventHandler(log logger.ILogger, ctx context.Context, redisClient *redis.Client) *KickPlayerRoomEventHandler {
	return &KickPlayerRoomEventHandler{
		log:         log,
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (h *KickPlayerRoomEventHandler) Handle(ctx context.Context, event *KickPlayerRoomEvent) error {
	roomId := event.RoomId
	playerId := event.PlayerId

	redisEvent := map[string]string{
		"room_id":   roomId,
		"player_id": playerId,
		"event":     "room:kicked",
	}
	data, err := json.Marshal(redisEvent)
	if err != nil {
		return err
	}

	if err := h.redisClient.Publish(h.ctx, "room_events", data).Err(); err != nil {
		return err
	}

	h.log.Info("KickPlayerRoomEventHandler invoked", "room_id", roomId, "player_id", playerId)
	return nil
}
