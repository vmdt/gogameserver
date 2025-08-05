package events

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/modules/boardgame/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type AttackBattleShipBoardEvent struct {
	PlayerId   string      `json:"player_id"`
	RoomId     string      `json:"room_id"`
	Shot       domain.Shot `json:"shot"`
	IsWin      bool        `json:"is_win" default:"false"`
	OpponentAt *time.Time  `json:"opponent_at,omitempty"`
}

func NewAttackBattleShipBoardEvent(playerId, roomId string, shot domain.Shot, opponentAt *time.Time, isWin bool) *AttackBattleShipBoardEvent {
	return &AttackBattleShipBoardEvent{
		PlayerId:   playerId,
		RoomId:     roomId,
		Shot:       shot,
		IsWin:      isWin,
		OpponentAt: opponentAt,
	}
}

type AttackBattleShipBoardEventHandler struct {
	log         logger.ILogger
	ctx         context.Context
	redisClient *redis.Client
}

func NewAttackBattleShipBoardEventHandler(log logger.ILogger, ctx context.Context, redisClient *redis.Client) *AttackBattleShipBoardEventHandler {
	return &AttackBattleShipBoardEventHandler{
		log:         log,
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (h *AttackBattleShipBoardEventHandler) Handle(ctx context.Context, event *AttackBattleShipBoardEvent) error {
	redisEvent := map[string]any{
		"player_id":   event.PlayerId,
		"room_id":     event.RoomId,
		"shot":        event.Shot,
		"opponent_at": event.OpponentAt,
		"event":       "battleship:attack",
	}

	data, err := json.Marshal(redisEvent)
	if err != nil {
		h.log.Error("Failed to marshal attack battleship event", "error", err)
		return err
	}

	if err := h.redisClient.Publish(h.ctx, "battleship_events", data).Err(); err != nil {
		h.log.Error("Failed to publish attack battleship event", "error", err)
		return err
	}

	return nil
}
