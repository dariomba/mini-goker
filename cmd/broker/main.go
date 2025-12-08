package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9092")
	if err != nil {
		panic(fmt.Sprintf("error creating tcp server: %s", err.Error()))
	}
	defer listener.Close()

	fmt.Println("tcp server started on port:", 9092)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection", err.Error())
			continue
		}

		go hanndleConnection(conn)
	}
}

func hanndleConnection(conn net.Conn) {
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

		fmt.Printf("received: %s", buffer[:n])
		conn.Write([]byte("message received\n"))
	}
}
