package query

import (
	"context"
	"github.com/mik3lon/go-htmx-starter/pkg/bus"
)

type QueryHandler interface {
	Handle(ctx context.Context, query bus.Dto) (interface{}, error)
}
