package client

import (
	"crypto/tls"
	"fmt"
)

type Client struct {
	Host string
	Port string
}

func NewClient(host, port string) *Client {
	return &Client{
		Host: host,
		Port: port,
	}
}

func (c *Client) Connect() (*tls.Conn, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // skip certificate verification
	}
	// connect
	conn, err := tls.Dial("tcp", c.Host+":"+c.Port, tlsConfig)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	fmt.Println("[debug] We did get here eventually")
	if err := VerifyServerCertificate(conn); err != nil {
		fmt.Println("[!] Closing the connection...")
		conn.Close()
		return nil, err
	}
	fmt.Println("[debug] Now it is time to open a shell to interact with the server")
	FileTransfer(conn)

	return conn, nil
}
