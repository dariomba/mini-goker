package main

import (
	"log"

	"github.com/dariomba/mini-goker/internal/routing"
	"github.com/dariomba/mini-goker/internal/transport"
)

func main() {
	//TODO: Load config from .yaml file
	listenAddr := ":9092"

	defaultHanddler := routing.NewDefaultHandler()

	server := transport.NewServer(listenAddr, defaultHanddler)

	log.Fatal(server.Start())
}
