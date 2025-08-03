package echo_middleware

import (
	"context"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/vmdt/gogameserver/pkg/auth"
)

func JwtAuthVerifier(ctx context.Context, jwtService auth.IJwtService, claimsType func() interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return echo.ErrUnauthorized
			}

			authHeaderParts := strings.Split(token, " ")
			if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
				return echo.ErrUnauthorized
			}

			verifiedToken, err := jwtService.Verify(authHeaderParts[1])
			if err != nil {
				return echo.ErrUnauthorized
			}

			claims := claimsType()
			if err := verifiedToken.Claims(claims); err != nil {
				return echo.ErrUnauthorized
			}

			c.Set("claims", claims)
			ctx = context.WithValue(ctx, "claims", claims)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
