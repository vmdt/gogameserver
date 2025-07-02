package echoserver

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/pkg/logger"
)

const (
	MaxHeaderBytes = 1 << 20
	ReadTimeout    = 15 * time.Second
	WriteTimeout   = 15 * time.Second
)

type EchoConfig struct {
	Port                string   `mapstructure:"port" validate:"required"`
	Development         bool     `mapstructure:"development"`
	BasePath            string   `mapstructure:"base_path" validate:"required"`
	DebugErrorsResponse bool     `mapstructure:"debug_errors_response"`
	IgnoreLogUrls       []string `mapstructure:"ignore_log_urls"`
	Timeout             int      `mapstructure:"timeout"`
	Host                string   `mapstructure:"host"`
}

func NewEchoServer() *echo.Echo {
	e := echo.New()
	return e
}

func RunHttpServer(ctx context.Context, echo *echo.Echo, log logger.ILogger, cfg *EchoConfig) error {
	echo.Server.ReadTimeout = ReadTimeout
	echo.Server.WriteTimeout = WriteTimeout
	echo.Server.MaxHeaderBytes = MaxHeaderBytes

	go func() {
		select {
		case <-ctx.Done():
			log.Infof("shutting down Http PORT: {%s}", cfg.Port)
			err := echo.Shutdown(ctx)
			if err != nil {
				log.Errorf("(Shutdown) err: {%v}", err)
				return
			}
			log.Info("server exited properly")
			return
		}
	}()

	err := echo.Start(cfg.Port)

	return err
}

func RegisterGroupFunc(groupName string, echo *echo.Echo, builder func(g *echo.Group)) *echo.Echo {
	builder(echo.Group(groupName))

	return echo
}
