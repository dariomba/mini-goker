package broker

import (
	"fmt"
	"io"
	"net"
)

type Server struct {
	listenAddr string
	listener   net.Listener
	quitch     chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
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

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("client disconnected")
				return
			}

			fmt.Println("error reading conn:", err)
			return
		}

		fmt.Printf("received: %s\n", buffer[:n])
		conn.Write([]byte("message received\n"))
	}
}
