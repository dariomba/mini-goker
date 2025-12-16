package transport

import (
	"context"
	"fmt"
	"net"

	"github.com/dariomba/mini-goker/internal/routing/protocol"
)

type Server struct {
	listenAddr string
	listener   net.Listener
	handler    Handler
	quitch     chan struct{}
}

func NewServer(listenAddr string, handler Handler) *Server {
	return &Server{
		listenAddr: listenAddr,
		handler:    handler,
		quitch:     make(chan struct{}),
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return fmt.Errorf("error creating tcp server: %s", err.Error())
	}
	defer listener.Close()
	s.listener = listener

	fmt.Println("tcp server started on", s.listenAddr)

	go s.acceptLoop()

	<-s.quitch

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection", err.Error())
			continue
		}

		fmt.Println("new connection to the server from:", conn.RemoteAddr().String())

		go s.hanndleConnection(conn)
	}
}

func (s *Server) hanndleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		req, err := protocol.DecodeFrame(conn)
		if err != nil {
			protocol.WriteError(conn, err)
			return
		}

		resp, err := s.handler.Handle(context.TODO(), req)
		if err != nil {
			protocol.WriteError(conn, err)
			continue
		}

		protocol.EncodeFrame(conn, resp)
	}
}
