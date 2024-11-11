package query

import (
	"context"
	"go-boilerplate/pkg/bus"
)

type QueryHandler interface {
	Handle(ctx context.Context, query bus.Dto) (interface{}, error)
}
