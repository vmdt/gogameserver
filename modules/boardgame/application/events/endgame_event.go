package events

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type EndgameEvent struct {
	Who      int    `json:"who"`
	PlayerId string `json:"player_id"`
	RoomId   string `json:"room_id"`
}

func NewEndgameEvent(who int, playerId, roomId string) *EndgameEvent {
	return &EndgameEvent{
		Who:      who,
		PlayerId: playerId,
		RoomId:   roomId,
	}
}

type EndgameEventHandler struct {
	log         logger.ILogger
	ctx         context.Context
	redisClient *redis.Client
}

func NewEndgameEventHandler(log logger.ILogger, ctx context.Context, redisClient *redis.Client) *EndgameEventHandler {
	return &EndgameEventHandler{
		log:         log,
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (h *EndgameEventHandler) Handle(ctx context.Context, event *EndgameEvent) error {
	redisEvent := map[string]any{
		"player_id": event.PlayerId,
		"room_id":   event.RoomId,
		"who_win":   event.Who,
		"event":     "room:endgame",
	}

	data, err := json.Marshal(redisEvent)
	if err != nil {
		h.log.Error("Failed to marshal attack battleship event", "error", err)
		return err
	}

	if err := h.redisClient.Publish(h.ctx, "room_events", data).Err(); err != nil {
		h.log.Error("Failed to publish room endgame event", "error", err)
		return err
	}

	return nil
}
