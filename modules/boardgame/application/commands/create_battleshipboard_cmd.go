package commands

import (
	"context"

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
	battleShip := &domain.BattleShip{
		PlayerId: uuid.MustParse(command.PlayerId),
		RoomId:   uuid.MustParse(command.RoomId),
		Ships:    command.Ships,
		Shots:    command.Shots,
	}
	createdBoard, err := h.bsRepo.CreateBoard(battleShip)
	if err != nil {
		return nil, err
	}
	return &dtos.BattleshipGame{
		PlayerId: createdBoard.PlayerId.String(),
		RoomId:   createdBoard.RoomId.String(),
		Ships:    createdBoard.Ships,
		Shots:    createdBoard.Shots,
	}, nil
}
