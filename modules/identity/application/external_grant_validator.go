package application

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vmdt/gogameserver/modules/identity/application/dtos"
	"github.com/vmdt/gogameserver/modules/identity/domain"
	"github.com/vmdt/gogameserver/pkg/auth"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type ExternalGrantValidator struct {
	log        logger.ILogger
	ctx        context.Context
	ssoService ISsoInfoService
	jwtService auth.IJwtService
}

func NewExternalGrantValidator(ctx context.Context, log logger.ILogger, ssoService ISsoInfoService, jwtService auth.IJwtService) *ExternalGrantValidator {
	return &ExternalGrantValidator{
		log:        log,
		ctx:        ctx,
		ssoService: ssoService,
		jwtService: jwtService,
	}
}

func (v *ExternalGrantValidator) Validate(ctx context.Context, email, provider string) (*dtos.UserAuthDTO, error) {
	var user *domain.User
	switch provider {
	case "google":
		userInfo, err := v.ssoService.GetExternalUserInfo(ctx, email, provider)
		if err != nil {
			v.log.Error("Failed to get external user info", "error", err)
			return nil, err
		}
		if userInfo == nil {
			v.log.Warn("No user info found for the provided email and provider")
			return nil, nil
		}

		user, _ = v.ssoService.GetSsoInfo(ctx, userInfo.Email, provider)
		if user.Email == "" {
			randomPass := uuid.New().String() // Generate a random password
			newUser := &domain.User{
				ID:       uuid.New(),
				Username: userInfo.Fullname,
				Email:    userInfo.Email,
				Provider: provider,
				Password: randomPass, // Set the random password
				IsSso:    true,
				Nation:   "Unknown",
			}

			err = v.ssoService.AddSsoUser(ctx, newUser)
			if err != nil {
				v.log.Error("Failed to add new SSO user", "error", err)
				return nil, err
			}
			user = newUser
		}
	default:
		return nil, nil
	}

	if user == nil {
		v.log.Warn("No SSO info found for the user", "email", email, "provider", provider)
		return nil, nil
	}

	claimAccessToken := map[string]interface{}{
		"user_id": user.ID,
	}

	claimRefreshToken := map[string]interface{}{
		"user_id": user.ID,
	}

	tokenPair, err := v.jwtService.CreateTokenPair(claimAccessToken, claimRefreshToken)
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
			ExpiresIn:    v.jwtService.GetTTL() / time.Second,
		},
	}, nil
}
