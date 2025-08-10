package main

import (
	"github.com/go-playground/validator"
	"github.com/vmdt/gogameserver/config"
	boardgame_api "github.com/vmdt/gogameserver/modules/boardgame/api"
	chat_api "github.com/vmdt/gogameserver/modules/chat/api"
	identity_api "github.com/vmdt/gogameserver/modules/identity/api"
	player_api "github.com/vmdt/gogameserver/modules/player/api"
	room_api "github.com/vmdt/gogameserver/modules/room/api"
	"github.com/vmdt/gogameserver/pkg/auth"
	echoserver "github.com/vmdt/gogameserver/pkg/echo"
	elastic "github.com/vmdt/gogameserver/pkg/elasticsearch"
	"github.com/vmdt/gogameserver/pkg/http"
	"github.com/vmdt/gogameserver/pkg/logger"
	"github.com/vmdt/gogameserver/pkg/postgresgorm"
	"github.com/vmdt/gogameserver/pkg/rabbitmq"
	redis2 "github.com/vmdt/gogameserver/pkg/redis"
	"github.com/vmdt/gogameserver/server"
	"github.com/vmdt/gogameserver/server/configurations"
	"go.uber.org/fx"

	_ "github.com/vmdt/gogameserver/pkg/system" // load system env
)

// @title Battleship API
// @version 1.0
// @description API for Battleship game
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and then your token.

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
// @description Provide your API key here.

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
				auth.NewJwtService,
				rabbitmq.NewRabbitMQConn,
				rabbitmq.NewPublisher,
			),
			player_api.Startup(),
			room_api.Startup(),
			boardgame_api.Startup(),
			identity_api.Startup(),
			chat_api.Startup(),
			fx.Invoke(server.RunAPIServer),
			fx.Invoke(configurations.ConfigSwagger),
			fx.Invoke(configurations.ConfigMiddleware),
		),
	).Run()
}
