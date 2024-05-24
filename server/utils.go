package server

import (
	"fmt"
	"io"
	"net"
	"os"
)

func GetFileFromClient(fileName string, conn net.Conn) error {
	defer conn.Close()
	// try n create a file on the server
	//file, err := os.Create(fileName)
	file, err := os.Create("new" + fileName) //! so i do not have problems with doubly named files...
	if err != nil {
		return fmt.Errorf("[!] Impossible to create the file: %v", err)
	}
	defer file.Close()

	n, err := io.Copy(file, conn)
	if err != nil {
		return fmt.Errorf("[!] Something went wrong while copying the file: %v", err)
	}

	fmt.Printf("[*] Received %d bytes and saved to %s\n", n, fileName)

	conn.Close()
	return nil
}
