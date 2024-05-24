package client

import (
	"bufio"
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
	fmt.Println("[*] Insert 'send <file>' if you want to send a file to the server, insert 'get <file>' to get a file from the server:")
	for {
		/*
			fmt.Scan(&userInput)
			fmt.Println("[debug] The user input is: " + userInput)
			//userInput = strings.TrimSpace(userInput) //! removes stuff which might be bad for the split, idk
		*/
		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		arrayUserInput := strings.Fields(userInput)
		fmt.Println("[debug] The array is split in: ", len(arrayUserInput))
		if len(arrayUserInput) < 2 {
			fmt.Println("[!] Insufficient arguments. Please use 'send <file>' or 'get <file>'.")
			continue
		}

		/*
			switch arrayUserInput[0] {
			case "send":
				fmt.Println("[*] Seems like you are trying to send a file to the server...")
				sendFilesToServer(arrayUserInput[1], conn)
			case "get":
				fmt.Println("[*] Seems like you are trying to get a file from the server...")
			default:
				fmt.Println("[!] Invalid syntax, insert 'send' or 'get' to send or get files to or from the server")
			}
		*/
		if arrayUserInput[0] == "send" {
			fmt.Println("[*] Seems like you are trying to send a file to the server...")
			err := sendFilesToServer(arrayUserInput[1], conn)
			if err != nil {
				fmt.Printf("[!] Error sending file: %s\n", err)
			}
		} else if arrayUserInput[0] == "get" {
			fmt.Println("[*] Seems like you are trying to get a file from the server...")
		} else {
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
		conn.Write([]byte("-1")) //shows to the server that you are a noob
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
