package main

import (
	"log"

	"github.com/dariomba/mini-goker/internal/broker"
)

func main() {
	listenAddr := ":9092"

	server := broker.NewServer(listenAddr)

	log.Fatal(server.Start())
}
