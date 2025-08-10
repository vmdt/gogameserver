package api

import (
	"github.com/vmdt/gogameserver/modules/chat/api/configurations"
	"github.com/vmdt/gogameserver/modules/chat/infrastructure"
	"go.uber.org/fx"
)

func Startup() fx.Option {
	return fx.Options(
		fx.Provide(infrastructure.NewChatDbContext),
		fx.Provide(infrastructure.NewChatRepositoryImp),
		fx.Invoke(configurations.ConfigMigrations),
		fx.Invoke(configurations.ConfigEndpoints),
		fx.Invoke(configurations.ConfigChatMediators),
	)
}
