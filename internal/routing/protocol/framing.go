package protocol

import "io"

func DecodeFrame(r io.Reader) (Request, error) {
	// read length
	// read type byte
	// decode payload into the correct struct
	return nil, nil
}

func WriteError(r io.Reader, err error) {
	// handle error
}

func EncodeFrame(r io.Reader, response Response) {
	// handle error
}
