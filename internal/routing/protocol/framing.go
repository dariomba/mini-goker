package protocol

import (
	"encoding/binary"
	"errors"
	"io"
)

const MaxFrameSize = 10 * 1024 * 1024 // 10 MB

func DecodeFrame(r io.Reader) (Request, error) {
	lenBuf := make([]byte, 4)
	_, err := io.ReadFull(r, lenBuf)
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lenBuf)
	if length == 0 || length > MaxFrameSize {
		return nil, errors.New("invalid frame length")
	}

	typeBuf := make([]byte, 1)
	_, err = io.ReadFull(r, typeBuf)
	if err != nil {
		return nil, err
	}
	msgType := MessageType(typeBuf[0])

	payloadBuf := make([]byte, length-1)
	_, err = io.ReadFull(r, payloadBuf)
	if err != nil {
		return nil, err
	}

	switch msgType {
	case MsgProduce:
		return &ProduceRequest{Payload: payloadBuf}, nil
	default:
		return &ProduceRequest{}, nil
	}
}

func WriteError(w io.Writer, err error) {
	// write dummy error response
	w.Write([]byte("error: " + err.Error() + "\n"))
}

func EncodeFrame(w io.Writer, response Response) {
	// write dummy response
	w.Write([]byte("message received\n"))
}
