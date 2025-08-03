package middlewares

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/pkg/auth"
	echo_middleware "github.com/vmdt/gogameserver/pkg/echo/middleware"
	"github.com/vmdt/gogameserver/server/pkg/jwt"
)

func JwtAuthVerifier(jwtService auth.IJwtService, ctx context.Context) echo.MiddlewareFunc {
	return echo_middleware.JwtAuthVerifier(ctx, jwtService, func() interface{} {
		return new(jwt.Claims)
	})
}
