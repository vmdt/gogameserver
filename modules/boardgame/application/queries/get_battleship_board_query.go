package queries

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/vmdt/gogameserver/modules/boardgame/application/dtos"
	"github.com/vmdt/gogameserver/modules/boardgame/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type GetBattleshipBoardQuery struct {
	PlayerId string `param:"player_id" validate:"required"`
	RoomId   string `param:"room_id" validate:"required"`
}

func NewGetBattleshipBoardQuery(playerId, roomId string) *GetBattleshipBoardQuery {
	return &GetBattleshipBoardQuery{
		PlayerId: playerId,
		RoomId:   roomId,
	}
}

type GetBattleshipBoardQueryHandler struct {
	log    logger.ILogger
	ctx    context.Context
	bsRepo domain.IBattleShipRepository
}

func NewGetBattleshipBoardQueryHandler(log logger.ILogger, ctx context.Context, bsRepo domain.IBattleShipRepository) *GetBattleshipBoardQueryHandler {
	return &GetBattleshipBoardQueryHandler{
		log:    log,
		ctx:    ctx,
		bsRepo: bsRepo,
	}
}

func (h *GetBattleshipBoardQueryHandler) Handle(ctx context.Context, query *GetBattleshipBoardQuery) (*dtos.BattleshipGame, error) {
	board, err := h.bsRepo.GetBoardGameByPlayerId(query.PlayerId, query.RoomId)
	if err != nil {
		h.log.Error("Failed to get battleship board", "error", err)
		return nil, err
	}
	if board == nil {
		h.log.Error("Battleship board not found", "player_id", query.PlayerId, "room_id", query.RoomId)
		return nil, errors.New("battleship board not found")
	}
	var ships []domain.Ship
	var shots []domain.Shot
	_ = json.Unmarshal(board.Ships, &ships)
	_ = json.Unmarshal(board.Shots, &shots)

	return &dtos.BattleshipGame{
		PlayerId: board.PlayerId.String(),
		RoomId:   board.RoomId.String(),
		Ships:    ships,
		Shots:    shots,
	}, nil
}
