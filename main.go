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
	flag.Usage = customUsage
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

func customUsage() {
	fmt.Fprintf(os.Stderr, "\n\nCadabra 1.0.0 (Git v1.0.0 packaged an 1.0.0-1)\nInteractively do something maybe\nSee https://github.com/a9sk/cadabra for more information.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: cdbr [options] ...\n\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	fmt.Fprintf(os.Stderr, "	-mode (client or server)\n")
	fmt.Fprintf(os.Stderr, "	-host (default: localhost)\n")
	fmt.Fprintf(os.Stderr, "	-port\n")
	fmt.Fprintf(os.Stderr, "	-file (path to the file if in client mode)\n\n")
	fmt.Fprintf(os.Stderr, "@a9sk\n")
}
