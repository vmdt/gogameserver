package main

import (
	"github.com/go-playground/validator"
	"github.com/vmdt/gogameserver/config"
	player_api "github.com/vmdt/gogameserver/modules/player/api"
	room_api "github.com/vmdt/gogameserver/modules/room/api"
	echoserver "github.com/vmdt/gogameserver/pkg/echo"
	"github.com/vmdt/gogameserver/pkg/http"
	"github.com/vmdt/gogameserver/pkg/logger"
	"github.com/vmdt/gogameserver/pkg/postgresgorm"
	redis2 "github.com/vmdt/gogameserver/pkg/redis"
	"github.com/vmdt/gogameserver/server"
	"github.com/vmdt/gogameserver/server/configurations"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.InitConfig,
				logger.InitLogger,
				http.NewContext,
				echoserver.NewEchoServer,
				postgresgorm.NewGorm,
				validator.New,
				redis2.NewRedisClient,
			),
			player_api.Startup(),
			room_api.Startup(),
			fx.Invoke(server.RunAPIServer),
			fx.Invoke(configurations.ConfigSwagger),
			fx.Invoke(configurations.ConfigMiddleware),
		),
	).Run()
}
