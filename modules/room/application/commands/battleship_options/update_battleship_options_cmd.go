package battleship_options_cmd

import (
	"context"

	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/room/application/dtos"
	"github.com/vmdt/gogameserver/modules/room/application/events"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type UpdateBattleshipOptionsCmd struct {
	TimePerTurn   *int   `json:"time_per_turn"`                         // in seconds
	TimePlaceShip *int   `json:"time_place_ship"`                       // in seconds
	WhoGoFirst    *int   `json:"who_go_first"`                          // 0: random
	RoomId        string `param:"room_id" validate:"required" json:"-"` // UUID of the room
}

func NewUpdateBattleshipOptionsCmd(timePerTurn, timePlaceShip, whoGoFirst *int, roomId string) *UpdateBattleshipOptionsCmd {
	return &UpdateBattleshipOptionsCmd{
		TimePerTurn:   timePerTurn,
		TimePlaceShip: timePlaceShip,
		WhoGoFirst:    whoGoFirst,
		RoomId:        roomId,
	}
}

type UpdateBattleshipOptionsCmdHandler struct {
	log logger.ILogger
	ctx context.Context
	db  *infrastructure.RoomDbContext
}

func NewUpdateBattleshipOptionsCmdHandler(log logger.ILogger, ctx context.Context, db *infrastructure.RoomDbContext) *UpdateBattleshipOptionsCmdHandler {
	return &UpdateBattleshipOptionsCmdHandler{
		log: log,
		ctx: ctx,
		db:  db,
	}
}

func (h *UpdateBattleshipOptionsCmdHandler) Handle(ctx context.Context, cmd *UpdateBattleshipOptionsCmd) (*dtos.BattleshipOptionsDTO, error) {
	var updates = map[string]interface{}{}
	if cmd.TimePerTurn != nil {
		updates["time_per_turn"] = *cmd.TimePerTurn
	}
	if cmd.TimePlaceShip != nil {
		updates["time_place_ship"] = *cmd.TimePlaceShip
	}
	if cmd.WhoGoFirst != nil {
		updates["who_go_first"] = *cmd.WhoGoFirst
	}
	if len(updates) == 0 {
		h.log.Warn("No updates provided for battleship options", "room_id", cmd.RoomId)
		return nil, nil
	}

	var battleshipOptions domain.BattleshipOptions
	dbResult := h.db.GetModelDB(&domain.BattleshipOptions{}).
		Where("room_id = ?", cmd.RoomId).
		Updates(updates).
		Scan(&battleshipOptions)
	if dbResult.Error != nil {
		h.log.Error("Failed to update and retrieve battleship options", "error", dbResult.Error)
		return nil, dbResult.Error
	}
	if dbResult.Error != nil {
		h.log.Error("Failed to update battleship options", "error", dbResult.Error)
		return nil, dbResult.Error
	}

	if cmd.WhoGoFirst != nil {
		var turn int
		if *cmd.WhoGoFirst == 0 {
			if uuid.New().ID()%2 == 0 {
				turn = 1
			} else {
				turn = 2
			}
		} else {
			turn = *cmd.WhoGoFirst
		}
		updateTurnEvent := &events.UpdateTurnEvent{
			RoomId: cmd.RoomId,
			Turn:   turn,
		}
		if err := mediatr.Publish(ctx, updateTurnEvent); err != nil {
			h.log.Error("Failed to publish update turn event", "error", err, "room_id", cmd.RoomId, "turn", turn)
			return nil, err
		}
	}

	updateOptsEvent := &events.UpdateBattleshipOptionsEvent{
		RoomId: cmd.RoomId,
	}
	if err := mediatr.Publish(ctx, updateOptsEvent); err != nil {
		h.log.Error("Failed to publish update battleship options event", "error", err, "room_id", cmd.RoomId)
		return nil, err
	}

	return &dtos.BattleshipOptionsDTO{
		Id:            battleshipOptions.ID.String(),
		TimePerTurn:   int(battleshipOptions.TimePerTurn),
		TimePlaceShip: int(battleshipOptions.TimePlaceShip),
		WhoGoFirst:    battleshipOptions.WhoGoFirst,
		RoomId:        cmd.RoomId,
	}, nil
}
