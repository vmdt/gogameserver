package api

import (
	"github.com/vmdt/gogameserver/modules/boardgame/api/configurations"
	"github.com/vmdt/gogameserver/modules/boardgame/infrastructure"
	"go.uber.org/fx"
)

func Startup() fx.Option {
	return fx.Options(
		fx.Provide(infrastructure.NewBoardGameDbContext),
		fx.Provide(infrastructure.NewBattleShipRepositoryImp),
		fx.Invoke(configurations.ConfigMigrations),
		fx.Invoke(configurations.ConfigBattleShipMediator),
		fx.Invoke(configurations.ConfigEndpoints),
	)
}
