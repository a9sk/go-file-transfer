package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func getFilesFromClient(fileName string, conn net.Conn) error {
	defer conn.Close()
	// try n create a file on the server
	file, err := os.Create("cdbr." + fileName)
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

func sendFilesToClient(fileName string, conn net.Conn) error {
	// imagine sending a non existing file...
	file, err := os.Open(strings.TrimSpace(fileName))
	if err != nil {
		conn.Write([]byte("[!] The file does not exist..."))
		return fmt.Errorf("[!] Could not open the file")
	}
	defer file.Close()

	_, err = conn.Write([]byte(fileName))
	if err != nil {
		return fmt.Errorf("[!] Could not write the file name")
	}

	n, err := io.Copy(conn, file)
	if err != nil {
		return fmt.Errorf("[!] Error while copying file contents: %v", err)
	}

	fmt.Printf("[debug] %d bytes sent", n)

	return nil
}
