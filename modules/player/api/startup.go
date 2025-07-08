package api

import (
	"github.com/vmdt/gogameserver/modules/player/api/configurations"
	"github.com/vmdt/gogameserver/modules/player/infrastructure"
	"go.uber.org/fx"
)

func Startup() fx.Option {
	return fx.Options(
		fx.Provide(infrastructure.NewPlayerDbContext),
		fx.Invoke(configurations.ConfigMigrations),
	)
}
