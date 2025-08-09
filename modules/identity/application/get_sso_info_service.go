package application

import (
	"context"
	"errors"

	"github.com/vmdt/gogameserver/modules/identity/domain"
	"github.com/vmdt/gogameserver/modules/identity/infrastructure"
	"github.com/vmdt/gogameserver/pkg/auth/google"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type (
	ExternalUserInfo struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		Fullname  string `json:"fullname"`
		Provider  string `json:"provider"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	ISsoInfoService interface {
		GetSsoInfo(ctx context.Context, email, provider string) (*domain.User, error)
		AddSsoUser(ctx context.Context, user *domain.User) error
		GetExternalUserInfo(ctx context.Context, token, provider string) (*ExternalUserInfo, error)
	}

	ssoInfoService struct {
		ctx          context.Context
		log          logger.ILogger
		dbContext    *infrastructure.IdentityDbContext
		googleClient *google.GoogleClient
	}
)

func NewSsoInfoService(ctx context.Context, log logger.ILogger, dbContext *infrastructure.IdentityDbContext) ISsoInfoService {
	return &ssoInfoService{
		ctx:          ctx,
		log:          log,
		dbContext:    dbContext,
		googleClient: google.NewGoogleClient(),
	}
}

func (s *ssoInfoService) GetSsoInfo(ctx context.Context, email, provider string) (*domain.User, error) {
	var user domain.User
	s.dbContext.GetModelDB(&domain.User{}).Where("email = ?", email).First(&user)
	return &user, nil
}

func (s *ssoInfoService) GetExternalUserInfo(ctx context.Context, token, provider string) (*ExternalUserInfo, error) {
	switch provider {
	case "google":
		userInfo, err := s.googleClient.GetUserInfo(token)
		if err != nil {
			return nil, err
		}
		return &ExternalUserInfo{
			ID:       userInfo.ID,
			Email:    userInfo.Email,
			Fullname: userInfo.Name,
			Provider: provider,
		}, nil
	}
	return nil, errors.New("unsupported provider")
}

func (s *ssoInfoService) AddSsoUser(ctx context.Context, user *domain.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	err := s.dbContext.GetModelDB(&domain.User{}).Create(user)
	if err.Error != nil {
		s.log.Error("Failed to add SSO user", "error", err)
		return err.Error
	}

	return nil
}
