package main

import (
	"log"

	"github.com/dariomba/mini-goker/internal/transport"
)

func main() {
	listenAddr := ":9092"

	server := transport.NewServer(listenAddr)

	log.Fatal(server.Start())
}
