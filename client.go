package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: client <server_address> <server_port>")
		os.Exit(1)
	}

	serverAddr := os.Args[1]
	serverPort := os.Args[2]

	for {
		fmt.Println("Enter a number: ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error Scanning: ", err)
			os.Exit(1)
		}

		conn, err := net.Dial("tcp", serverAddr+":"+serverPort)
		if err != nil {
			fmt.Println("Error connecting to the server:", err)
			continue
		}
	}
}
