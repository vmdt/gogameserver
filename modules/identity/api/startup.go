package api

import (
	"github.com/vmdt/gogameserver/modules/identity/api/configurations"
	"github.com/vmdt/gogameserver/modules/identity/infrastructure"
	"github.com/vmdt/gogameserver/pkg/auth"
	"go.uber.org/fx"
)

func Startup() fx.Option {
	return fx.Options(
		fx.Provide(infrastructure.NewIdentityDbContext),
		fx.Provide(infrastructure.NewUserRepositoryImp),
		fx.Provide(auth.NewJwtService),
		fx.Invoke(configurations.ConfigMigrations),
		fx.Invoke(configurations.ConfigEndpoints),
		fx.Invoke(configurations.ConfigMediators),
	)
}
