package queries

import (
	"context"
	"encoding/json"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/boardgame/application/dtos"
	"github.com/vmdt/gogameserver/modules/boardgame/application/events"
	"github.com/vmdt/gogameserver/modules/boardgame/domain"
	"github.com/vmdt/gogameserver/modules/boardgame/infrastructure"
	room_domain "github.com/vmdt/gogameserver/modules/room/domain"
	room_infrastructure "github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type CheckSunkShipStatusQuery struct {
	RoomId   string `param:"room_id" validate:"required"`
	PlayerId string `param:"player_id" validate:"required"`
}

func NewCheckSunkShipStatusQuery(roomId, playerId string) *CheckSunkShipStatusQuery {
	return &CheckSunkShipStatusQuery{
		RoomId:   roomId,
		PlayerId: playerId,
	}
}

type CheckSunkShipStatusQueryHandler struct {
	log           logger.ILogger
	ctx           context.Context
	dbContext     *infrastructure.BoardGameDbContext
	roomDbContext *room_infrastructure.RoomDbContext
}

func NewCheckSunkShipStatusQueryHandler(log logger.ILogger, ctx context.Context, dbContext *infrastructure.BoardGameDbContext, roomDbContext *room_infrastructure.RoomDbContext) *CheckSunkShipStatusQueryHandler {
	return &CheckSunkShipStatusQueryHandler{
		log:           log,
		ctx:           ctx,
		dbContext:     dbContext,
		roomDbContext: roomDbContext,
	}
}

func (h *CheckSunkShipStatusQueryHandler) Handle(ctx context.Context, query *CheckSunkShipStatusQuery) (*dtos.SunkShipsDTO, error) {
	myBoard := &domain.BattleShip{}
	if err := h.dbContext.GetModelDB(&domain.BattleShip{}).Where("player_id = ? AND room_id = ?", query.PlayerId, query.RoomId).First(myBoard).Error; err != nil {
		h.log.Error("Failed to get battleship board", "error", err, "player_id", query.PlayerId, "room_id", query.RoomId)
		return nil, err
	}

	myPlayer := &room_domain.RoomPlayer{}
	if err := h.roomDbContext.GetModelDB(&room_domain.RoomPlayer{}).Where("player_id = ? AND room_id = ?", query.PlayerId, query.RoomId).First(myPlayer).Error; err != nil {
		h.log.Error("Failed to get room player", "error", err, "player_id", query.PlayerId, "room_id", query.RoomId)
		return nil, err
	}

	oppBoard := &domain.BattleShip{}
	if err := h.dbContext.GetModelDB(&domain.BattleShip{}).Where("player_id != ? AND room_id = ?", query.PlayerId, query.RoomId).First(oppBoard).Error; err != nil {
		h.log.Error("Failed to get opponent's battleship board", "error", err, "player_id", query.PlayerId, "room_id", query.RoomId)
		return nil, err
	}

	var myShots []domain.Shot
	var oppShips []domain.Ship
	_ = json.Unmarshal(myBoard.Shots, &myShots)
	_ = json.Unmarshal(oppBoard.Ships, &oppShips)

	isSamePosition := func(a, b domain.Position) bool {
		return a.X == b.X && a.Y == b.Y
	}

	var sunkShips []dtos.SunkShipDTO

	numShipSunk := 0
	for _, ship := range oppShips {
		sunk := true

		for _, pos := range ship.Positions {
			hit := false
			for _, shot := range myShots {
				if isSamePosition(shot.Position, pos) && shot.Status == "hit" {
					hit = true
					break
				}
			}
			if !hit {
				sunk = false
				break
			}
		}
		if sunk {
			numShipSunk++
		}

		sunkShips = append(sunkShips, dtos.SunkShipDTO{
			ShipName: ship.Name,
			Size:     ship.Size,
			IsSunk:   sunk,
		})
	}

	if numShipSunk == 5 {
		endgameEvent := events.NewEndgameEvent(myPlayer.Me, query.PlayerId, query.RoomId)
		if err := mediatr.Publish(ctx, endgameEvent); err != nil {
			h.log.Error("Failed to publish endgame event", "error", err)
			return nil, err
		}
	}

	return &dtos.SunkShipsDTO{
		PlayerId:  query.PlayerId,
		Ships:     sunkShips,
		NumOfSunk: numShipSunk,
	}, nil
}
