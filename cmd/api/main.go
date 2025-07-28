package main

import (
	"github.com/go-playground/validator"
	"github.com/vmdt/gogameserver/config"
	boardgame_api "github.com/vmdt/gogameserver/modules/boardgame/api"
	player_api "github.com/vmdt/gogameserver/modules/player/api"
	room_api "github.com/vmdt/gogameserver/modules/room/api"
	echoserver "github.com/vmdt/gogameserver/pkg/echo"
	elastic "github.com/vmdt/gogameserver/pkg/elasticsearch"
	"github.com/vmdt/gogameserver/pkg/http"
	"github.com/vmdt/gogameserver/pkg/logger"
	"github.com/vmdt/gogameserver/pkg/postgresgorm"
	redis2 "github.com/vmdt/gogameserver/pkg/redis"
	"github.com/vmdt/gogameserver/server"
	"github.com/vmdt/gogameserver/server/configurations"
	"go.uber.org/fx"

	_ "github.com/vmdt/gogameserver/pkg/system" // load system env
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
				elastic.NewElasticClient,
			),
			player_api.Startup(),
			room_api.Startup(),
			boardgame_api.Startup(),
			fx.Invoke(server.RunAPIServer),
			fx.Invoke(configurations.ConfigSwagger),
			fx.Invoke(configurations.ConfigMiddleware),
		),
	).Run()
}
