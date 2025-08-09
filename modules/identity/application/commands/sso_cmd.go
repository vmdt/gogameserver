package commands

import (
	"context"

	"github.com/vmdt/gogameserver/modules/identity/application"
	"github.com/vmdt/gogameserver/modules/identity/application/dtos"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type ExternalGrantCommand struct {
	Credential string `json:"credential" validate:"required"`
	Provider   string `json:"-"`
}

func NewExternalGrantCommand(credential, provider string) *ExternalGrantCommand {
	return &ExternalGrantCommand{
		Credential: credential,
		Provider:   provider,
	}
}

type ExternalGrantCommandHandler struct {
	log            logger.ILogger
	ctx            context.Context
	grantValidator *application.ExternalGrantValidator
}

func NewExternalGrantCommandHandler(log logger.ILogger, ctx context.Context, grantValidator *application.ExternalGrantValidator) *ExternalGrantCommandHandler {
	return &ExternalGrantCommandHandler{
		log:            log,
		ctx:            ctx,
		grantValidator: grantValidator,
	}
}

func (h *ExternalGrantCommandHandler) Handle(ctx context.Context, command *ExternalGrantCommand) (*dtos.UserAuthDTO, error) {
	result, err := h.grantValidator.Validate(ctx, command.Credential, command.Provider)
	if err != nil {
		h.log.Error("ExternalGrantCommandHandler: Failed to validate external grant", "error", err)
		return nil, err
	}

	if result == nil {
		h.log.Warn("ExternalGrantCommandHandler: No user found for the provided credential and provider")
		return nil, nil
	}

	return result, nil
}
