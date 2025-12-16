package main

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/dariomba/mini-goker/internal/routing/protocol"
)

func main() {
	connectAddr := ":9092"

	conn, err := net.Dial("tcp", connectAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	payload := []byte(`{"topic":"orders","message":"hello world"}`)

	length := uint32(1 + len(payload))

	frame := make([]byte, 4+1+len(payload))

	binary.BigEndian.PutUint32(frame[0:4], length)

	// Message type (PRODUCE = 1)
	frame[4] = byte(protocol.MsgProduce)

	copy(frame[5:], payload)

	_, err = conn.Write(frame)
	if err != nil {
		panic(err)
	}

	// Read response
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Print("Response from server: ", string(buf[:n]))
}
