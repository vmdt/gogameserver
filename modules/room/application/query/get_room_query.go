package query

import (
	"context"
	"errors"

	"github.com/vmdt/gogameserver/modules/room/application/dtos"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type GetRoomQuery struct {
	ID string `json:"id"`
}

func NewGetRoomQuery(id string) *GetRoomQuery {
	return &GetRoomQuery{
		ID: id,
	}
}

type GetRoomHandler struct {
	log       logger.ILogger
	ctx       context.Context
	roomRepo  domain.IRoomRepository
	dbContext *infrastructure.RoomDbContext
}

func NewGetRoomHandler(log logger.ILogger, ctx context.Context, roomRepo domain.IRoomRepository, dbContext *infrastructure.RoomDbContext) *GetRoomHandler {
	return &GetRoomHandler{
		log:       log,
		ctx:       ctx,
		roomRepo:  roomRepo,
		dbContext: dbContext,
	}
}

func (h *GetRoomHandler) Handle(ctx context.Context, query *GetRoomQuery) (*dtos.RoomInformationDTO, error) {
	room, err := h.roomRepo.GetRoomByID(ctx, query.ID)
	if err != nil {
		h.log.Error("Failed to fetch room", "error", err)
		return nil, errors.New("room not found")
	}

	var battleshipOptions domain.BattleshipOptions
	errBattleshipOptions := h.dbContext.GetModelDB(&domain.BattleshipOptions{}).Where("room_id = ?", room.ID).First(&battleshipOptions).Error
	if errBattleshipOptions != nil {
		h.log.Error("GetRoomHandler: Failed to fetch battleship options", "error", errBattleshipOptions)
	}

	var battleshipOptionsDTO *dtos.BattleshipOptionsDTO
	if errBattleshipOptions == nil {
		battleshipOptionsDTO = &dtos.BattleshipOptionsDTO{
			Id:            battleshipOptions.ID.String(),
			TimePerTurn:   int(battleshipOptions.TimePerTurn),
			TimePlaceShip: int(battleshipOptions.TimePlaceShip),
			WhoGoFirst:    battleshipOptions.WhoGoFirst,
			StartPlaceAt:  battleshipOptions.StartPlaceAt,
			RoomId:        room.ID.String(),
		}
	} else {
		battleshipOptionsDTO = nil
	}

	roomDto := &dtos.RoomDTO{
		ID:                room.ID.String(),
		Status:            room.Status,
		Turn:              room.Turn,
		BattleshipOptions: battleshipOptionsDTO,
		CreatedAt:         room.CreatedAt,
		UpdatedAt:         room.UpdatedAt,
	}

	var roomPlayers []*domain.RoomPlayer
	if err := h.dbContext.GetModelDB(&domain.RoomPlayer{}).Where("room_id = ?", room.ID).Preload("Player").Find(&roomPlayers).Error; err != nil {
		h.log.Error("Failed to fetch room players", "error", err)
		return nil, err
	}
	var roomPlayerDtos []*dtos.RoomPlayerDTO
	for _, rp := range roomPlayers {
		roomPlayerDtos = append(roomPlayerDtos, rp.ToDTO())
	}

	h.log.Info("GetRoomQuery: Room fetched successfully", "room_id ", room.ID, "players_count ", len(roomPlayerDtos))

	return &dtos.RoomInformationDTO{
		Room:    *roomDto,
		Players: roomPlayerDtos,
	}, nil
}
