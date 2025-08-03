package jwt

import "github.com/kataras/jwt"

type Claims struct {
	jwt.Claims
	Role                string `json:"role"`
	UserId              string `json:"user_id"`
	TwoFactorVerified   bool   `json:"two_factor_verified"`
	IsSubscriptionValid bool   `json:"is_subscription_valid"`
}
