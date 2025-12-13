package routing

import (
	"context"
	"fmt"

	"github.com/dariomba/mini-goker/internal/routing/protocol"
)

type DefaultHandler struct {
}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}

func (h *DefaultHandler) Handle(ctx context.Context, req protocol.Request) (protocol.Response, error) {
	reqType := req.Type()
	switch reqType {
	case protocol.MsgProduce:
		fmt.Println("handling produce request...")
	default:
		return nil, fmt.Errorf("unknown request type")
	}

	return &protocol.ProduceResponse{}, nil
}
