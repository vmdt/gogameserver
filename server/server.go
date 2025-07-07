package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/config"
	echoserver "github.com/vmdt/gogameserver/pkg/echo"
	"github.com/vmdt/gogameserver/pkg/logger"
	"go.uber.org/fx"
)

func RunAPIServer(
	lc fx.Lifecycle,
	log logger.ILogger,
	e *echo.Echo,
	ctx context.Context,
	cfg *config.Config,
) error {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := echoserver.RunHttpServer(ctx, e, log, cfg.Echo); !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("error running http server: %v", err)
				}
			}()

			// Health check endpoint
			e.GET("/", func(c echo.Context) error {
				return c.String(http.StatusOK, "ok")
			})

			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Info("shutting down server")
			return nil
		},
	})

	return nil
}
