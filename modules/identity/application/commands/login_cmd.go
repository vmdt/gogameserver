package commands

import (
	"context"
	"errors"
	"time"

	"github.com/vmdt/gogameserver/modules/identity/application/dtos"
	"github.com/vmdt/gogameserver/modules/identity/domain"
	"github.com/vmdt/gogameserver/pkg/auth"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type LoginCommand struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

func NewLoginCommand(email, password string) *LoginCommand {
	return &LoginCommand{
		Email:    email,
		Password: password,
	}
}

type LoginCommandHandler struct {
	log        logger.ILogger
	ctx        context.Context
	userRepo   domain.IUserRepository
	jwtService auth.IJwtService
}

func NewLoginCommandHandler(log logger.ILogger, ctx context.Context, userRepo domain.IUserRepository, jwtService auth.IJwtService) *LoginCommandHandler {
	return &LoginCommandHandler{
		log:        log,
		ctx:        ctx,
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (h *LoginCommandHandler) Handle(ctx context.Context, command *LoginCommand) (*dtos.UserAuthDTO, error) {
	user, err := h.userRepo.GetUserByEmail(command.Email)
	if err != nil {
		h.log.Error("LoginCommandHandler: Failed to get user by email", "error", err)
		return nil, err
	}
	if user == nil {
		h.log.Error("LoginCommandHandler: User not found", "email", command.Email)
		return nil, errors.New("user not found")
	}

	matchedPassword, err := user.ValidatePassword(command.Password)
	if err != nil {
		h.log.Error("LoginCommandHandler: Failed to validate password", "error", err)
		return nil, err
	}
	if !matchedPassword {
		h.log.Error("LoginCommandHandler: Invalid email or password", "email", command.Email)
		return nil, errors.New("invalid email or password")
	}

	claimAccessToken := map[string]interface{}{
		"user_id": user.ID,
	}
	claimRefreshToken := map[string]interface{}{
		"user_id": user.ID,
	}

	tokenPair, err := h.jwtService.CreateTokenPair(claimAccessToken, claimRefreshToken)
	if err != nil {
		return nil, err
	}

	return &dtos.UserAuthDTO{
		User: &dtos.UserDTO{
			ID:        user.ID.String(),
			Username:  user.Username,
			Email:     user.Email,
			Nation:    user.Nation,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		TokenPair: &dtos.TokenPairDTO{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
			ExpiresIn:    h.jwtService.GetTTL() / time.Second,
		},
	}, nil
}
