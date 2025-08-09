package configurations

import (
	"context"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/redis/go-redis/v9"
	"github.com/vmdt/gogameserver/modules/identity/application"
	"github.com/vmdt/gogameserver/modules/identity/application/commands"
	"github.com/vmdt/gogameserver/modules/identity/domain"
	"github.com/vmdt/gogameserver/modules/identity/infrastructure"
	"github.com/vmdt/gogameserver/pkg/auth"
	"github.com/vmdt/gogameserver/pkg/logger"
)

func ConfigMediators(
	log logger.ILogger,
	ctx context.Context,
	userRepo domain.IUserRepository,
	jwtService auth.IJwtService,
	db *infrastructure.IdentityDbContext,
	redisClient *redis.Client,
	ssoService application.ISsoInfoService,
) error {
	// Register command handlers
	err := mediatr.RegisterRequestHandler(
		commands.NewRegisterUserCommandHandler(log, ctx, userRepo, jwtService),
	)
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler(
		commands.NewLoginCommandHandler(log, ctx, userRepo, jwtService),
	)
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler(
		commands.NewRefreshTokenCommandHandler(log, ctx, userRepo, jwtService),
	)
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler(
		commands.NewExternalGrantCommandHandler(log, ctx, application.NewExternalGrantValidator(ctx, log, ssoService, jwtService)),
	)
	if err != nil {
		return err
	}

	return nil
}
