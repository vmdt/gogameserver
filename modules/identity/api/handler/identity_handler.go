package handler

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/vmdt/gogameserver/modules/identity/application/commands"
	"github.com/vmdt/gogameserver/modules/identity/application/dtos"
)

// RegisterUserHandler handles user registration requests
// @Summary      Register a new user
// @Description  Registers a new user and returns the authentication token.
// @Tags         Identity
// @Accept       json
// @Produce      json
// @Param        body         body      commands.RegisterUserCommand     true   "User registration details"
// @Success      200          {object}  dtos.UserAuthDTO
// @Failure      400          {object}  map[string]string
// @Failure      500          {object}  map[string]string
// @Router       /api/v1/identity/register [post]
func Register(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req commands.RegisterUserCommand
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}
		if err := validator.Struct(req); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}

		result, err := mediatr.Send[*commands.RegisterUserCommand, *dtos.UserAuthDTO](c.Request().Context(), &req)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, result)
	}
}

// LoginHandler handles user login requests
// @Summary      User login
// @Description  Authenticates a user and returns the authentication token.
// @Tags         Identity
// @Accept       json
// @Produce      json
// @Param        body         body      commands.LoginCommand     true   "User login details"
// @Success      200          {object}  dtos.UserAuthDTO
// @Failure      400          {object}  map[string]string
// @Failure      500          {object}  map[string]string
// @Router       /api/v1/identity/login [post]
func Login(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req commands.LoginCommand
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}
		if err := validator.Struct(req); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}

		result, err := mediatr.Send[*commands.LoginCommand, *dtos.UserAuthDTO](c.Request().Context(), &req)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, result)
	}
}

// RefreshTokenHandler handles refresh token requests
// @Summary      Refresh user token
// @Description  Refreshes the user's authentication token using a refresh token.
// @Tags         Identity
// @Accept       json
// @Produce      json
// @Param        body         body      commands.RefreshTokenCommand     true   "Refresh token details"
// @Success      200          {object}  dtos.TokenPairDTO
// @Failure      400          {object}  map[string]string
// @Failure      500          {object}  map[string]string
// @Router       /api/v1/identity/refresh [post]
func RefreshToken(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req commands.RefreshTokenCommand
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}
		if err := validator.Struct(req); err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}

		result, err := mediatr.Send[*commands.RefreshTokenCommand, *dtos.TokenPairDTO](c.Request().Context(), &req)
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, result)
	}
}
