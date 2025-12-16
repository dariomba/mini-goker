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
		return h.handleProduce(ctx, req)
	default:
		return nil, fmt.Errorf("unknown request type")
	}
}

func (h *DefaultHandler) handleProduce(ctx context.Context, req protocol.Request) (protocol.Response, error) {
	produceReq, ok := req.(*protocol.ProduceRequest)
	if !ok {
		return nil, fmt.Errorf("invalid produce request")
	}

	fmt.Println("Received produce request with payload:", string(produceReq.Payload))

	offset := int64(0) // Dummy offset for illustration

	resp := &protocol.ProduceResponse{
		Offset: offset,
		Err:    "",
	}

	return resp, nil
}
