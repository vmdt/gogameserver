package battleship_options_cmd

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/vmdt/gogameserver/modules/room/application/dtos"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type CreateBattleshipOptionsCmd struct {
	TimePerTurn   int    `json:"time_per_turn"`               // in seconds
	TimePlaceShip int    `json:"time_place_ship"`             // in seconds
	WhoGoFirst    int    `json:"who_go_first"`                // 0: random
	RoomId        string `json:"room_id" validate:"required"` // UUID of the room
}

func NewCreateBattleshipOptionsCmd(timePerTurn, timePlaceShip, whoGoFirst int, roomId string) *CreateBattleshipOptionsCmd {
	return &CreateBattleshipOptionsCmd{
		TimePerTurn:   timePerTurn,
		TimePlaceShip: timePlaceShip,
		WhoGoFirst:    whoGoFirst,
		RoomId:        roomId,
	}
}

type CreateBattleshipOptionsCmdHandler struct {
	log logger.ILogger
	ctx context.Context
	db  *infrastructure.RoomDbContext
}

func NewCreateBattleshipOptionsCmdHandler(log logger.ILogger, ctx context.Context, db *infrastructure.RoomDbContext) *CreateBattleshipOptionsCmdHandler {
	return &CreateBattleshipOptionsCmdHandler{
		log: log,
		ctx: ctx,
		db:  db,
	}
}

func (h *CreateBattleshipOptionsCmdHandler) Handle(ctx context.Context, cmd *CreateBattleshipOptionsCmd) (*dtos.BattleshipOptionsDTO, error) {
	battleshipOptions := &domain.BattleshipOptions{
		ID:            uuid.New(),
		TimePerTurn:   time.Duration(cmd.TimePerTurn),
		TimePlaceShip: time.Duration(cmd.TimePlaceShip),
		WhoGoFirst:    cmd.WhoGoFirst,
		RoomId:        uuid.MustParse(cmd.RoomId),
	}

	if err := h.db.GetModelDB(&domain.BattleshipOptions{}).Create(battleshipOptions).Error; err != nil {
		h.log.Error("Failed to create battleship options", "error", errors.Wrap(err, "CreateBattleshipOptionsCmdHandler.Handle"))
		return nil, errors.Wrap(err, "failed to create battleship options")
	}

	h.log.Info("Battleship options created successfully", "room_id", cmd.RoomId)
	return &dtos.BattleshipOptionsDTO{
		Id:            battleshipOptions.ID.String(),
		TimePerTurn:   cmd.TimePerTurn,
		TimePlaceShip: cmd.TimePlaceShip,
		WhoGoFirst:    cmd.WhoGoFirst,
		RoomId:        cmd.RoomId,
	}, nil
}
