package api

import (
	"github.com/vmdt/gogameserver/modules/player/api/configurations"
	"github.com/vmdt/gogameserver/modules/player/infrastructure"
	"go.uber.org/fx"
)

func Startup() fx.Option {
	return fx.Options(
		fx.Provide(infrastructure.NewPlayerDbContext),
		fx.Provide(infrastructure.NewPlayerRepository),
		fx.Invoke(configurations.ConfigMigrations),
		fx.Invoke(configurations.ConfigPlayerMediator),
		fx.Invoke(configurations.ConfigEndpoints),
	)
}
