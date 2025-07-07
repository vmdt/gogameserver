package configurations

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/room/application/commands"
	"github.com/vmdt/gogameserver/modules/room/application/query"
	"github.com/vmdt/gogameserver/modules/room/domain"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigRoomMediator(log logger.ILogger, ctx context.Context, roomRepo domain.IRoomRepository) error {
	err := mediatr.RegisterRequestHandler(commands.NewCreateRoomHandler(log, ctx, roomRepo))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler(query.NewGetRoomHandler(log, ctx, roomRepo))
	if err != nil {
		return err
	}

	log.Info("Room mediator configurations completed successfully")
	return nil
}
