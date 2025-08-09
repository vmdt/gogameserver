package api

import (
	"github.com/vmdt/gogameserver/modules/identity/api/configurations"
	"github.com/vmdt/gogameserver/modules/identity/application"
	"github.com/vmdt/gogameserver/modules/identity/infrastructure"
	"go.uber.org/fx"
)

func Startup() fx.Option {
	return fx.Options(
		fx.Provide(infrastructure.NewIdentityDbContext),
		fx.Provide(infrastructure.NewUserRepositoryImp),
		fx.Provide(application.NewSsoInfoService),
		fx.Invoke(configurations.ConfigMigrations),
		fx.Invoke(configurations.ConfigEndpoints),
		fx.Invoke(configurations.ConfigMediators),
	)
}
