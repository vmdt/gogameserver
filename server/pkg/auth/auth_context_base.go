package auth

import (
	"context"

	"github.com/vmdt/gogameserver/server/pkg/jwt"
)

func GetUserId(ctx context.Context) string {
	claims, ok := ctx.Value("claims").(*jwt.Claims)
	if !ok || claims == nil {
		return ""
	}
	return claims.UserId
}
