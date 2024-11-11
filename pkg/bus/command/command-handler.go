package command

import (
	"context"
	"go-boilerplate/pkg/bus"
)

type CommandHandler interface {
	Handle(ctx context.Context, command bus.Dto) error
}
