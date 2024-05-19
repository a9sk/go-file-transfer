// server/server.go

package server

import (
	"crypto/tls"
	"fmt"
	"net"
)

type Server struct {
	Port string
	// fields and configs
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) ListenAndServe() error {
	// use TLS certificates
	cert, err := tls.LoadX509KeyPair("certificates/server.crt", "certificates/server.key")
	if err != nil {
		return fmt.Errorf("error loading certificates: %v", err)
	}

	// create TLS config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// shhhhhhh, listen
	listener, err := tls.Listen("tcp", "localhost:"+s.Port, tlsConfig) //! CHANGE THE localhost: to only : WHEN PUBLISHING
	if err != nil {
		return fmt.Errorf("error starting server: %v", err)
	}
	defer listener.Close()

	fmt.Println("[*] Server started. Listening on port", s.Port)

	// let others connect
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[!] Error accepting connection:", err)
			continue
		}

		// handle connections
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	fmt.Println("[debug] Someone connected")
	defer conn.Close()
	// add logic for receiving files, handling requests, etc.
}

// add methods for handling connections, receiving files, etc.
