package client

import (
	"crypto/tls"
	"fmt"
)

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
