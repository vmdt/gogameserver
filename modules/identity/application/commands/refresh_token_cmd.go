package commands

import (
	"context"
	"errors"
	"time"

	"github.com/vmdt/gogameserver/modules/identity/application/dtos"
	"github.com/vmdt/gogameserver/modules/identity/domain"
	"github.com/vmdt/gogameserver/pkg/auth"
	"github.com/vmdt/gogameserver/pkg/logger"
	"github.com/vmdt/gogameserver/server/pkg/jwt"
)

type RefreshTokenCommand struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func NewRefreshTokenCommand(refreshToken string) *RefreshTokenCommand {
	return &RefreshTokenCommand{
		RefreshToken: refreshToken,
	}
}

type RefreshTokenCommandHandler struct {
	log        logger.ILogger
	ctx        context.Context
	userRepo   domain.IUserRepository
	jwtService auth.IJwtService
}

func NewRefreshTokenCommandHandler(log logger.ILogger, ctx context.Context, userRepo domain.IUserRepository, jwtService auth.IJwtService) *RefreshTokenCommandHandler {
	return &RefreshTokenCommandHandler{
		log:        log,
		ctx:        ctx,
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (h *RefreshTokenCommandHandler) Handle(ctx context.Context, command *RefreshTokenCommand) (*dtos.TokenPairDTO, error) {
	verifiedToken, err := h.jwtService.Verify(command.RefreshToken)
	if err != nil {
		h.log.Error("Failed to verify refresh token", "error", err)
		return nil, err
	}
	claims := &jwt.Claims{}
	if err := verifiedToken.Claims(claims); err != nil {
		h.log.Error("Failed to parse claims from refresh token", "error", err)
		return nil, errors.New("invalid token")
	}

	userId := claims.UserId
	if userId == "" {
		h.log.Error("User ID not found in claims")
		return nil, errors.New("invalid token")
	}

	user, err := h.userRepo.GetUserById(userId)
	if err != nil {
		h.log.Error("Failed to get user by ID", "error", err, "user_id", userId)
		return nil, errors.New("invalid token")
	}
	if user == nil {
		h.log.Error("User not found", "user_id", userId)
		return nil, errors.New("user not found")
	}
	claimAccessToken := map[string]interface{}{
		"user_id": user.ID,
	}
	claimRefreshToken := map[string]interface{}{
		"user_id": user.ID,
	}
	tokenPair, err := h.jwtService.CreateTokenPair(claimAccessToken, claimRefreshToken)
	if err != nil {
		h.log.Error("Failed to create token pair", "error", err)
		return nil, err
	}

	return &dtos.TokenPairDTO{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresIn:    h.jwtService.GetTTL() / time.Second,
	}, nil
}
