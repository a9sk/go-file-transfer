package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"os"
	"strings"
)

const BUFFER_SIZE = 1024 //! might want to change the buffer size...

func VerifyServerCertificate(conn *tls.Conn) error {
	fmt.Println("[debug] We also got here ggs")
	// check if certificate is provided
	if len(conn.ConnectionState().PeerCertificates) == 0 {
		return fmt.Errorf("server did not provide a certificate")
	}

	if len(conn.ConnectionState().VerifiedChains) == 0 || len(conn.ConnectionState().VerifiedChains[0]) == 0 || !conn.ConnectionState().VerifiedChains[0][0].IsCA {
		fmt.Println("[!] Server's certificate is not signed by a trusted CA.")
		fmt.Println("[?] Do you want to proceed and trust the server? (yes/no): ")
		var resp string

		for {
			fmt.Scanln(&resp)
			switch resp {
			case "yes":
				fmt.Println("[*] Proceeding with the connection...")
				return nil
			case "no":
				return fmt.Errorf("connection aborted by user")
			default:
				fmt.Println("[*] Please enter 'yes' or 'no'")
			}
		}
	}
	fmt.Println("[*] Best case, Server's certificate is signed by a trusted CA ;)")

	return nil
}

func FileTransfer(conn *tls.Conn) error {
	fmt.Println("[debug] FileTransfer is cool func")
	fmt.Println("[*] Insert 'send' if you want to send a file to the server, insert 'get' to get a file from the server:")
	var userInput string
	for {
		fmt.Scanln(&userInput)
		arrayUserInput := strings.Split(userInput, " ")
		switch arrayUserInput[0] {
		case "send":
			fmt.Println("[*] Seems like you are trying to send a file to the server...")
			sendFilesToServer(arrayUserInput[1], conn)
		case "get":
			fmt.Println("[*] Seems like you are trying to get a file from the server...")
		default:
			fmt.Println("[!] Invalid syntax, insert 'send' or 'get' to send or get files to or from the server")
		}
	}
}

func sendFilesToServer(fileName string, conn *tls.Conn) error {
	defer conn.Close()

	fileBuffer := make([]byte, BUFFER_SIZE)

	// imagine sending a non existing file...
	_, err := os.Stat(strings.TrimSpace(fileName))
	if os.IsNotExist(err) {
		conn.Write([]byte("-1"))
		return fmt.Errorf("[!] The file does not exist")
	}

	file, err := os.Open(strings.TrimSpace(fileName))
	if err != nil {
		conn.Write([]byte("-1"))
		return fmt.Errorf("[!] Could not open the file")
	}
	defer file.Close()

	_, err = conn.Write([]byte("send " + fileName))
	if err != nil {
		return fmt.Errorf("[!] Could not write the connection bytes")
	}

	// read the file buffer and send it until it "exists"
	for {
		n, err := file.Read(fileBuffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("[!] Cannot read the filebufffer")
		}
		if n == 0 {
			break
		}
		_, err = conn.Write(fileBuffer[:n])
		if err != nil {
			return fmt.Errorf("[!] Impossible to write the filebuffer")
		}
	}
	return nil
}
