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
	// add check to see if a service is already on on this port...
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
	//listener, err := tls.Listen("tcp", "192.168.1.9:"+s.Port, tlsConfig) //! CHANGE THE host to only : WHEN PUBLISHING
	listener, err := tls.Listen("tcp", "localhost:"+s.Port, tlsConfig)
	//listener, err := tls.Listen("tcp", ":"+s.Port, tlsConfig)
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
		fmt.Println("[debug] someone connected")
		// handle connections
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("[debug] handling connection")

	// loop to read data
	for {
		buffer := make([]byte, 1024) //create a buffer max size... //! Pay attention
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("[!] Error reading from connection:", err)
			return
		}

		// process the data...
		fmt.Println("[*] Received data from client:", string(buffer))
	}
}

// add methods for handling connections, receiving files, etc.
