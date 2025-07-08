package api

import (
	"github.com/vmdt/gogameserver/modules/room/api/configurations"
	"github.com/vmdt/gogameserver/modules/room/infrastructure"
	"go.uber.org/fx"
)

func Startup() fx.Option {
	return fx.Options(
		fx.Provide(infrastructure.NewRoomDbContext),
		fx.Provide(infrastructure.NewRoomRepository),
		fx.Invoke(configurations.ConfigMigrations),
		fx.Invoke(configurations.ConfigRoomMediator),
		fx.Invoke(configurations.ConfigEndpoints),
	)
}
