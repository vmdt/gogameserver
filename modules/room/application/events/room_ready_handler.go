package events

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type RoomReadyEvent struct {
	RoomId   string `json:"room_id"`
	PlayerId string `json:"player_id"`
}

func (e *RoomReadyEvent) GetRoomId() string {
	return e.RoomId
}

func (e *RoomReadyEvent) GetPlayerId() string {
	return e.PlayerId
}

type RoomReadyHandler[T IRoomReadyEvent] struct {
	log         logger.ILogger
	ctx         context.Context
	db          *infrastructure.RoomDbContext
	redisClient *redis.Client
}

func NewRoomReadyHandler[T IRoomReadyEvent](log logger.ILogger, ctx context.Context, redisClient *redis.Client, db *infrastructure.RoomDbContext) *RoomReadyHandler[T] {
	return &RoomReadyHandler[T]{
		log:         log,
		ctx:         ctx,
		redisClient: redisClient,
		db:          db,
	}
}

func (h *RoomReadyHandler[T]) Handle(ctx context.Context, event T) error {
	roomId := event.GetRoomId()
	playerId := event.GetPlayerId()

	redisEvent := map[string]string{
		"room_id":   roomId,
		"player_id": playerId,
		"event":     "room:started",
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
