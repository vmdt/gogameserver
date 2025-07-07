package main

import (
	"github.com/vmdt/gogameserver/config"
	room_api "github.com/vmdt/gogameserver/modules/room/api"
	echoserver "github.com/vmdt/gogameserver/pkg/echo"
	"github.com/vmdt/gogameserver/pkg/http"
	"github.com/vmdt/gogameserver/pkg/logger"
	"github.com/vmdt/gogameserver/pkg/postgresgorm"
	"github.com/vmdt/gogameserver/server"
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
			),
			fx.Invoke(server.RunAPIServer),
			room_api.Startup(),
		),
	).Run()
}
