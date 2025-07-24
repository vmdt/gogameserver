package commands

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/boardgame/domain"
	room_dtos "github.com/vmdt/gogameserver/modules/room/application/dtos"
	room_query "github.com/vmdt/gogameserver/modules/room/application/query"
	"github.com/vmdt/gogameserver/pkg/logger"
	"github.com/vmdt/gogameserver/pkg/utils"
)

type AttackBattleShipCommand struct {
	PlayerId string          `json:"player_id"`
	RoomId   string          `json:"room_id"`
	Position domain.Position `json:"position"`
}

func NewAttackBattleShipCommand(playerId, roomId string, position domain.Position) *AttackBattleShipCommand {
	return &AttackBattleShipCommand{
		PlayerId: playerId,
		RoomId:   roomId,
		Position: position,
	}
}

type AttackBattleShipCommandHandler struct {
	log    logger.ILogger
	ctx    context.Context
	bsRepo domain.IBattleShipRepository
}

func NewAttackBattleShipCommandHandler(log logger.ILogger, ctx context.Context, bsRepo domain.IBattleShipRepository) *AttackBattleShipCommandHandler {
	return &AttackBattleShipCommandHandler{
		log:    log,
		ctx:    ctx,
		bsRepo: bsRepo,
	}
}

func (h *AttackBattleShipCommandHandler) Handle(ctx context.Context, command *AttackBattleShipCommand) (bool, error) {
	board, err := h.bsRepo.GetBoardGameByPlayerId(command.PlayerId, command.RoomId)
	if err != nil {
		h.log.Error("Failed to get battleship board", "error", err)
		return false, err
	}
	if board == nil {
		h.log.Error("Battleship board not found", "player_id", command.PlayerId, "room_id", command.RoomId)
		return false, errors.New("battleship board not found")
	}

	roomQuery := room_query.NewGetRoomQuery(command.RoomId)
	roomPlayers, err := mediatr.Send[*room_query.GetRoomQuery, *room_dtos.RoomInformationDTO](ctx, roomQuery)
	var shotStatus string = "miss"
	var oppShips []domain.Ship
	if len(roomPlayers.Players) > 1 {
		oppPlayer := utils.Filter(roomPlayers.Players, func(p *room_dtos.RoomPlayerDTO) bool {
			return p.PlayerId != command.PlayerId
		})[0]
		opponent, err := h.bsRepo.GetBoardGameByPlayerId(oppPlayer.PlayerId, command.RoomId)
		if err != nil {
			h.log.Error("Failed to get opponent's battleship board", "error", err)
			return false, err
		}
		_ = json.Unmarshal(opponent.Ships, &oppShips)

		// check if the shot hit an opponent's ship
		for _, ship := range oppShips {
			for _, shipPosition := range ship.Positions {
				if shipPosition.X == command.Position.X && shipPosition.Y == command.Position.Y {
					h.log.Info("Hit opponent's ship", "player_id", command.PlayerId, "room_id", command.RoomId)
					shotStatus = "hit"
				}
			}
		}
	}
	// Update the shots for the current player
	var shots []domain.Shot
	_ = json.Unmarshal(board.Shots, &shots)
	shots = append(shots, domain.Shot{
		Position: command.Position,
		Status:   shotStatus,
	})
	board.Shots, err = json.Marshal(shots)
	_, err = h.bsRepo.AddOrUpdate(board)
	if err != nil {
		h.log.Error("Failed to update battleship board", "error", err)
		return false, err
	}

	if shotStatus == "miss" {
		return false, nil
	}

	return true, nil
}
