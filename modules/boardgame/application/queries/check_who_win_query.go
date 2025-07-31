package queries

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/vmdt/gogameserver/modules/boardgame/application/dtos"
	"github.com/vmdt/gogameserver/modules/boardgame/domain"
	"github.com/vmdt/gogameserver/modules/boardgame/infrastructure"
	room_domain "github.com/vmdt/gogameserver/modules/room/domain"
	room_infrastructure "github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type CheckWhoWinQuery struct {
	RoomId string `param:"room_id" validate:"required"`
}

func NewCheckWhoWinQuery(roomId string) *CheckWhoWinQuery {
	return &CheckWhoWinQuery{
		RoomId: roomId,
	}
}

type CheckWhoWinQueryHandler struct {
	log           logger.ILogger
	ctx           context.Context
	dbContext     *infrastructure.BoardGameDbContext
	roomDbContext *room_infrastructure.RoomDbContext
}

func NewCheckWhoWinQueryHandler(log logger.ILogger, ctx context.Context, dbContext *infrastructure.BoardGameDbContext, roomDbContext *room_infrastructure.RoomDbContext) *CheckWhoWinQueryHandler {
	return &CheckWhoWinQueryHandler{
		log:           log,
		ctx:           ctx,
		dbContext:     dbContext,
		roomDbContext: roomDbContext,
	}
}

func (h *CheckWhoWinQueryHandler) Handle(ctx context.Context, query *CheckWhoWinQuery) (*dtos.WhoWinDTO, error) {
	room := &room_domain.Room{}
	if err := h.roomDbContext.GetModelDB(&room_domain.Room{}).Where("id = ?", query.RoomId).First(room).Error; err != nil {
		h.log.Error("Failed to get room by ID", "error", err, "room_id", query.RoomId)
		return nil, err
	}

	if room == nil {
		h.log.Error("Room not found", "room_id", query.RoomId)
		return nil, errors.New("room not found")
	}

	roomPlayers := []room_domain.RoomPlayer{}
	if err := h.roomDbContext.GetModelDB(&room_domain.RoomPlayer{}).Where("room_id = ?", query.RoomId).Find(&roomPlayers).Error; err != nil {
		h.log.Error("Failed to get room players", "error", err, "room_id", query.RoomId)
		return nil, err
	}
	if len(roomPlayers) < 2 {
		h.log.Error("Not enough players in the room", "room_id", query.RoomId)
		return nil, errors.New("not enough players in the room")
	}

	whoWinDTO := &dtos.WhoWinDTO{
		RoomId:    query.RoomId,
		WinStatus: []dtos.WinStatusDTO{},
	}
	for _, player := range roomPlayers {
		whoWinDTO.WinStatus = append(whoWinDTO.WinStatus, dtos.WinStatusDTO{
			PlayerId: player.PlayerId.String(),
			Win:      false, // default
			Placed:   false, // default
		})
	}

	bsBoards := []domain.BattleShip{}
	if err := h.dbContext.GetModelDB(&domain.BattleShip{}).Where("room_id = ?", query.RoomId).Find(&bsBoards).Error; err != nil {
		h.log.Error("Failed to get battleship boards", "error", err, "room_id", query.RoomId)
		return nil, err
	}

	if room.Status == "setup" {
		for _, board := range bsBoards {
			var ships []domain.Ship
			if err := json.Unmarshal(board.Ships, &ships); err != nil {
				h.log.Error("Failed to unmarshal ships", "error", err)
				return nil, err
			}

			if len(ships) == 5 {
				for i, status := range whoWinDTO.WinStatus {
					if status.PlayerId == board.PlayerId.String() {
						whoWinDTO.WinStatus[i].Placed = true
						break
					}
				}
			}
		}
	} else if room.Status == "battle" {
		// TODO: Implement logic to check who wins based on shots
	}

	return whoWinDTO, nil
}
