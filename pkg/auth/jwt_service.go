package auth

import (
	"strconv"
	"time"

	"github.com/vmdt/gogameserver/pkg/system"

	"github.com/kataras/jwt"
)

type (
	IJwtService interface {
		CreateTokenPair(claim interface{}, refreshTokenClaim interface{}) (jwt.TokenPair, error)
		CreateToken(claim interface{}) (string, error)
		Verify(token string) (*jwt.VerifiedToken, error)
		GetTTL() time.Duration
	}
	jwtService struct {
	}
)

func NewJwtService() IJwtService {
	return &jwtService{}
}

func (s *jwtService) GetTTL() time.Duration {
	ttl, _ := strconv.Atoi(system.Getenv("AUTH_JWT_TTL", "60"))
	return time.Duration(ttl) * time.Minute
}

func (s *jwtService) CreateTokenPair(claim interface{}, refreshTokenClaim interface{}) (jwt.TokenPair, error) {

	secret := system.Getenv("AUTH_JWT_SECRET")
	if secret == "" {
		panic("Please setup SecretKey")
	}

	refreshTokenDuration, _ := strconv.Atoi(system.Getenv("AUTH_JWT_FRESH_TTL", "43200"))

	accessToken, err := jwt.Sign(jwt.HS256, []byte(secret), claim, jwt.MaxAge(s.GetTTL()))
	if err != nil {
		return jwt.TokenPair{}, err
	}

	refreshToken, err := jwt.Sign(jwt.HS256, []byte(secret), refreshTokenClaim, jwt.MaxAge(time.Duration(refreshTokenDuration)*time.Minute))
	if err != nil {
		return jwt.TokenPair{}, err
	}
	return jwt.NewTokenPair(accessToken, refreshToken), nil
}

func (s *jwtService) CreateToken(claim interface{}) (string, error) {

	secret := system.Getenv("AUTH_JWT_SECRET")
	if secret == "" {
		panic("Please setup SecretKey")
	}

	accessToken, err := jwt.Sign(jwt.HS256, []byte(secret), claim, jwt.MaxAge(s.GetTTL()))
	if err != nil {
		return "", err
	}

	return string(accessToken), nil
}

func (s *jwtService) Verify(token string) (*jwt.VerifiedToken, error) {
	secret := system.Getenv("AUTH_JWT_SECRET")
	if secret == "" {
		panic("Please setup SecretKey")
	}
	return jwt.Verify(jwt.HS256, []byte(secret), []byte(token), jwt.Plain)

}
