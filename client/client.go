// client/client.go

package client

import (
	"crypto/tls"
	"fmt"
)

type Client struct {
	Host string
	Port string
	// client-specific stuff and configurations should go here
}

func NewClient(host, port string) *Client {
	return &Client{
		Host: host,
		Port: port,
	}
}

func (c *Client) Connect() (*tls.Conn, error) {
	// connect
	conn, err := tls.Dial("tcp", c.Host+":"+c.Port, &tls.Config{})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return conn, nil
}

// other methods for sending files, handling responses, etc.
