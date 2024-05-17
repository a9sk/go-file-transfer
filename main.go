// main.go

package main

import (
	"flag"
	"fmt"
	"os"

	"cadabra/client"
	"cadabra/server"
)

func main() {
	// command-line flags
	mode := flag.String("mode", "", "Mode: 'client' or 'server'")
	host := flag.String("host", "localhost", "Host address")
	port := flag.String("port", "8080", "Port number")
	filePath := flag.String("file", "", "File path (client mode)")
	flag.Parse()

	// validation
	if *mode != "client" && *mode != "server" {
		fmt.Println("Error: Mode must be 'client' or 'server'")
		flag.Usage()
		os.Exit(1)
	}

	switch *mode {
	case "client":
		if *filePath == "" {
			fmt.Println("Error: File path is required in client mode")
			flag.Usage()
			os.Exit(1)
		}

		// create and connect client
		client := client.NewClient(*host, *port)
		conn, err := client.Connect()
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			os.Exit(1)
		}
		defer conn.Close()

	case "server":
		// create and start server
		server := server.NewServer(*port)
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
		}
	}
}
