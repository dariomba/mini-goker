package transport

import (
	"context"

	"github.com/dariomba/mini-goker/internal/routing/protocol"
)

type Handler interface {
	Handle(ctx context.Context, req protocol.Request) (protocol.Response, error)
}
