package client

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"os"
	"strings"
)

const BUFFER_SIZE = 4000 //! might want to change the buffer size...

func VerifyServerCertificate(conn *tls.Conn) error {
	//fmt.Println("[debug] We also got here ggs")
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
	//fmt.Println("[debug] FileTransfer is cool func")
	fmt.Println("[*] Insert 'send <file>' if you want to send a file to the server, insert 'get <file>' to get a file from the server:")
	for {
		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		arrayUserInput := strings.Fields(userInput)
		//fmt.Println("[debug] The array is split in: ", len(arrayUserInput))
		if len(arrayUserInput) < 2 {
			fmt.Println("[!] Insufficient arguments. Please use 'send <file>' or 'get <file>'.")
			continue
		}
		if arrayUserInput[0] == "send" {
			fmt.Println("[*] Seems like you are trying to send a file to the server...")
			err := sendFilesToServer(arrayUserInput[1], conn)
			if err != nil {
				return fmt.Errorf("[!] Error sending file: %s", err)
			}
			break
		} else if arrayUserInput[0] == "get" {
			fmt.Println("[*] Seems like you are trying to get a file from the server...")
			err := getFilesFromServer(arrayUserInput[1], conn)
			if err != nil {
				return fmt.Errorf("[!] Error getting file: %s", err)
			}
			break
		} else {
			fmt.Println("[!] Invalid syntax, insert 'send' or 'get' to send or get files to or from the server")
		}
	}
	return nil
}

func sendFilesToServer(fileName string, conn *tls.Conn) error {
	// imagine sending a non existing file...
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

	// to copy file to the connection w-out buffer strange stuff
	n, err := io.Copy(conn, file)
	if err != nil {
		return fmt.Errorf("[!] Error while copying file contents: %v", err)
	}

	fmt.Printf("[*] %d bytes sent", n)

	return nil
}

func getFilesFromServer(fileName string, conn *tls.Conn) error {
	conn.Write([]byte("get " + fileName))

	// this legit is the same code the server uses...

	buffer := make([]byte, BUFFER_SIZE)
	l, err := conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("[!] Error reading from connection: %v", err)
	}
	//fmt.Println("[debug] Received data from server:", string(buffer))
	recivedName := string(buffer[:l])

	if fileName == recivedName {
		//fmt.Println("[debug] filename sent correctly")
		file, err := os.Create("cdbr." + fileName) // to recognize them
		if err != nil {
			return fmt.Errorf("[!] Impossible to create the file: %v", err)
		}
		defer file.Close()
		//fmt.Println("[debug] 1")
		n, err := io.Copy(file, conn)
		if err != nil {
			return fmt.Errorf("[!] Something went wrong while copying the file: %v", err)
		}
		//fmt.Print("[debug] 2")
		fmt.Printf("[*] Received %d bytes and saved to %s", n, fileName)

		conn.Close()
	} else {
		return fmt.Errorf("[!] The server returned a different file from the one you asked: %s", recivedName)
	}

	return nil
}
