package commands

import (
	"context"
	"time"

	"github.com/vmdt/gogameserver/modules/identity/application/dtos"
	"github.com/vmdt/gogameserver/modules/identity/domain"
	"github.com/vmdt/gogameserver/pkg/auth"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type RegisterUserCommand struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
	Nation   string `json:"nation"`
}

func NewRegisterUserCommand(username, email, password, nation string) *RegisterUserCommand {
	return &RegisterUserCommand{
		Username: username,
		Email:    email,
		Password: password,
		Nation:   nation,
	}
}

type RegisterUserCommandHandler struct {
	log        logger.ILogger
	ctx        context.Context
	userRepo   domain.IUserRepository
	jwtService auth.IJwtService
}

func NewRegisterUserCommandHandler(log logger.ILogger, ctx context.Context, userRepo domain.IUserRepository, jwtService auth.IJwtService) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{
		log:        log,
		ctx:        ctx,
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (h *RegisterUserCommandHandler) Handle(ctx context.Context, command *RegisterUserCommand) (*dtos.UserAuthDTO, error) {
	user := &domain.User{
		Username: command.Username,
		Email:    command.Email,
		Password: command.Password,
		Nation:   command.Nation,
	}

	createdUser, err := h.userRepo.CreateUser(user)
	if err != nil {
		h.log.Error("RegisterUserCommandHandler: Failed to create user", "error", err)
		return nil, err
	}
	claimAccessToken := map[string]interface{}{
		"user_id": createdUser.ID,
	}
	claimRefreshToken := map[string]interface{}{
		"user_id": createdUser.ID,
	}

	tokenPair, err := h.jwtService.CreateTokenPair(claimAccessToken, claimRefreshToken)
	if err != nil {
		return nil, err
	}

	return &dtos.UserAuthDTO{
		User: &dtos.UserDTO{
			ID:        createdUser.ID.String(),
			Username:  createdUser.Username,
			Email:     createdUser.Email,
			Nation:    createdUser.Nation,
			CreatedAt: createdUser.CreatedAt,
			UpdatedAt: createdUser.UpdatedAt,
		},
		TokenPair: &dtos.TokenPairDTO{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
			ExpiresIn:    h.jwtService.GetTTL() / time.Second,
		},
	}, nil
}
