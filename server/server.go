package server

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"
)

const BUFFER_SIZE = 4000

type Server struct {
	Port string
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
	listener, err := tls.Listen("tcp", ":"+s.Port, tlsConfig)
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
		//fmt.Println("[debug] someone connected")
		// handle connections
		go s.handleConnection(conn)
		// should use a whole own thread if i understood it right
	}
}

func (s *Server) handleConnection(conn net.Conn) error {
	defer conn.Close()
	//fmt.Println("[debug] handling connection")

	// loop to read data
	for {
		buffer := make([]byte, BUFFER_SIZE)
		l, err := conn.Read(buffer)
		if err != nil {
			return fmt.Errorf("[!] Error reading from connection: %v", err)
		}

		receivedMessage := string(buffer[:l])
		fields := strings.Fields(receivedMessage)
		if len(fields) < 2 {
			return fmt.Errorf("[!] There should only be 2 fields")
		}

		command := fields[0]
		fileName := fields[1]

		switch command {
		case "send":
			getFilesFromClient(fileName, conn)
			//fmt.Println("[debug] Received data from client:", string(buffer))
		case "get":
			sendFilesToClient(fileName, conn)
		default:
			return fmt.Errorf("[???] How did you even manage to get this error???")
		}

		conn.Close()
	}
}
