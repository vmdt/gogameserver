package commands

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/vmdt/gogameserver/modules/boardgame/application/dtos"
	"github.com/vmdt/gogameserver/modules/boardgame/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type CreateBattleShipBoardCommand struct {
	PlayerId string        `json:"player_id"`
	RoomId   string        `json:"room_id"`
	Ships    []domain.Ship `json:"ships"`
	Shots    []domain.Shot `json:"shots"`
}

func NewCreateBattleShipBoardCommand(
	playerId, roomId string, ships []domain.Ship, shots []domain.Shot,
) *CreateBattleShipBoardCommand {
	return &CreateBattleShipBoardCommand{
		PlayerId: playerId,
		RoomId:   roomId,
		Ships:    ships,
		Shots:    shots,
	}
}

type CreateBattleShipBoardCommandHandler struct {
	log    logger.ILogger
	ctx    context.Context
	bsRepo domain.IBattleShipRepository
}

func NewCreateBattleShipBoardCommandHandler(log logger.ILogger, ctx context.Context, bsRepo domain.IBattleShipRepository) *CreateBattleShipBoardCommandHandler {
	return &CreateBattleShipBoardCommandHandler{
		log:    log,
		ctx:    ctx,
		bsRepo: bsRepo,
	}
}

func (h *CreateBattleShipBoardCommandHandler) Handle(ctx context.Context, command *CreateBattleShipBoardCommand) (*dtos.BattleshipGame, error) {
	shipsJson, err := json.Marshal(command.Ships)
	if err != nil {
		h.log.Error("Failed to marshal ships", "error", err)
		return nil, err
	}
	shotsJson, err := json.Marshal(command.Shots)
	if err != nil {
		h.log.Error("Failed to marshal shots", "error", err)
		return nil, err
	}
	battleShip := &domain.BattleShip{
		PlayerId: uuid.MustParse(command.PlayerId),
		RoomId:   uuid.MustParse(command.RoomId),
		Ships:    shipsJson,
		Shots:    shotsJson,
	}
	createdBoard, err := h.bsRepo.AddOrUpdate(battleShip)
	if err != nil {
		return nil, err
	}

	var ships []domain.Ship
	var shots []domain.Shot
	_ = json.Unmarshal(createdBoard.Ships, &ships)
	_ = json.Unmarshal(createdBoard.Shots, &shots)

	return &dtos.BattleshipGame{
		PlayerId: createdBoard.PlayerId.String(),
		RoomId:   createdBoard.RoomId.String(),
		Ships:    ships,
		Shots:    shots,
	}, nil
}
