package configurations

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/redis/go-redis/v9"
	boardgame_events "github.com/vmdt/gogameserver/modules/boardgame/application/events"
	"github.com/vmdt/gogameserver/modules/room/application/commands"
	player_room_cmd "github.com/vmdt/gogameserver/modules/room/application/commands/player_room"
	"github.com/vmdt/gogameserver/modules/room/application/events"
	"github.com/vmdt/gogameserver/modules/room/application/query"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigRoomMediator(
	log logger.ILogger,
	ctx context.Context,
	roomRepo domain.IRoomRepository,
	db *infrastructure.RoomDbContext,
	redisClient *redis.Client,
) error {
	err := mediatr.RegisterRequestHandler(commands.NewCreateRoomHandler(log, ctx, roomRepo))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler(commands.NewPlayerCreateRoomHandler(log, ctx, roomRepo))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler(commands.NewJoinRoomHandler(log, ctx, roomRepo, db))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler(player_room_cmd.NewUpdateRoomPlayerCommandHandler(log, ctx, db))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler(player_room_cmd.NewKickPlayerRoomCommandHandler(log, ctx, db))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler(commands.NewUpdateRoomStatusCommandHandler(log, ctx, roomRepo))
	if err != nil {
		return err
	}

	// Register internal command handler for creating room players
	err = mediatr.RegisterRequestHandler(player_room_cmd.NewInternalCreateRoomPlayerCommandHandler(log, ctx, db))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler(query.NewGetRoomHandler(log, ctx, roomRepo, db))
	if err != nil {
		return err
	}

	// Register events
	err = mediatr.RegisterNotificationHandler(events.NewJoinRoomEventHandler(log, ctx, redisClient))
	if err != nil {
		return err
	}

	err = mediatr.RegisterNotificationHandler(events.NewKickPlayerRoomEventHandler(log, ctx, redisClient))
	if err != nil {
		return err
	}

	err = mediatr.RegisterNotificationHandler(events.NewRoomReadyHandler[*boardgame_events.ReadyBattleShipBoardEvent](log, ctx, redisClient, db))
	if err != nil {
		return err
	}

	err = mediatr.RegisterNotificationHandler(events.NewUpdateRoomPlayerStatusHandler[*boardgame_events.UpdatePlayerStatusEvent](log, ctx, redisClient, db))
	if err != nil {
		return err
	}

	err = mediatr.RegisterNotificationHandler(events.NewAttackBattleShipBoardEventHandler(log, ctx, db, redisClient))
	if err != nil {
		return err
	}

	log.Info("Room mediator configurations completed successfully")
	return nil
}
